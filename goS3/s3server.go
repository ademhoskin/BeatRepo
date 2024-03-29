package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateS3Client() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
		return nil
	}
	return s3.NewFromConfig(cfg)
}

// s3UploadFile uploads a file to an S3 bucket.
// It takes the bucket name, object key, and file path as input parameters.
// Returns an error if any error occurs during the upload process.
func S3UploadFile(bucketName, objectKey, filePath string) error {
	client := CreateS3Client()
	uploader := manager.NewUploader(client)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

// s3DownloadFile downloads a file from an S3 bucket and saves it to the specified file path.
// It takes the bucket name, object key, and file path as input parameters.
// Returns an error if any error occurs during the download process.
func S3DownloadFile(bucketName, objectKey, filePath string) error {
	client := CreateS3Client()
	downloader := manager.NewDownloader(client)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	if err != nil {
		return err
	}

	return nil
}

// s3DeleteFile deletes a file from an S3 bucket.
// It takes the bucket name and object key as parameters and returns an error if any.
func S3DeleteFile(bucketName, objectKey string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	if err != nil {
		return err
	}

	return nil
}

// handlers
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := S3UploadFile(data["bucketName"], data["objectKey"], data["filePath"]); err != nil {
		log.Fatalf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bucketName := r.URL.Query().Get("bucketName")
	objectKey := r.URL.Query().Get("objectKey")
	filePath := r.URL.Query().Get("filePath")

	S3DownloadFile(bucketName, objectKey, filePath)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bucketName := r.URL.Query().Get("bucketName")
	objectKey := r.URL.Query().Get("objectKey")

	if err := S3DeleteFile(bucketName, objectKey); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// server configuration
func ExecServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", UploadHandler)
	mux.HandleFunc("/download", DownloadHandler)
	mux.HandleFunc("/delete", DeleteHandler)
	http.ListenAndServe(":8080", mux)
	fmt.Println("Server started at port 8080")
}

// run the server
func main() {
	ExecServer()
}
