package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"net/http"
	"os"
	"testing"
)

func TestS3UploadDownloadIntegration(t *testing.T) {
	// Setup
	bucketName := "cppbeatproj"
	objectKey := "test-object"
	uploadFilePath := "test-upload.txt"
	downloadFilePath := "test-downloaded.txt"

	CreateTestFile(uploadFilePath, t)
	defer os.Remove(uploadFilePath)

	// Upload the file to S3
	if err := S3UploadFile(bucketName, objectKey, uploadFilePath); err != nil {
		t.Fatalf("Failed to upload test file to S3: %v", err)
	}

	//defer the deletion of the test object from S3
	defer func() {
		if err := S3DeleteFile(bucketName, objectKey); err != nil {
			t.Fatalf("Failed to delete test object from S3: %v", err)
		}
	}()

	// Download the file from S3
	if err := S3DownloadFile(bucketName, objectKey, downloadFilePath); err != nil {
		t.Fatalf("Failed to download test file from S3: %v", err)
	}
	defer os.Remove(downloadFilePath)

	// Compare the uploaded and downloaded files
	if !FilesAreIdentical(uploadFilePath, downloadFilePath) {
		t.Fatalf("Uploaded and downloaded files are not identical (did you mess around with the file content?)")
	}
}

func TestHandlers(t *testing.T) {
	// Setup
	bucketName := "cppbeatproj"
	objectKey := "test-object"
	uploadFilePath := "test-upload.txt"
	downloadFilePath := "test-downloaded.txt"
	s3Client := CreateS3Client()

	// Start the server
	ExecServer()

	// test upload handler

	CreateTestFile(uploadFilePath, t)
	defer os.Remove(uploadFilePath)

	uploadURL := "http://localhost:8080/upload"
	reqBody := CreateRequestBody(bucketName, objectKey, uploadFilePath)
	res, err := http.Post(uploadURL, "application/json", reqBody)
	if err != nil {
		t.Fatalf("Failed to send POST request to upload handler: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Upload handler returned status code %d", res.StatusCode)
	}

	// test download handler
	downloadURL := "http://localhost:8080/download"
	req, err := http.NewRequest(http.MethodGet, downloadURL, nil)
	if err != nil {
		t.Fatalf("Failed to create GET request to download handler: %v", err)
	}

	q := req.URL.Query()
	q.Add("bucketName", bucketName)
	q.Add("objectKey", objectKey)
	q.Add("filePath", downloadFilePath)
	req.URL.RawQuery = q.Encode()

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send GET request to download handler: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Download handler returned status code %d", res.StatusCode)
	}
	// Compare the uploaded and downloaded files
	if !FilesAreIdentical(uploadFilePath, downloadFilePath) {
		t.Fatalf("Uploaded and downloaded files are not identical (did you mess around with the file content?)")
	}

	// test delete handler
	deleteURL := "http://localhost:8080/delete"
	reqBody = CreateRequestBody(bucketName, objectKey, "")
	req, err = http.NewRequest(http.MethodDelete, deleteURL, reqBody)
	if err != nil {
		t.Fatalf("Failed to create DELETE request to delete handler: %v", err)
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send DELETE request to delete handler: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Delete handler returned status code %d", res.StatusCode)
	}

	// Check if the object was delete
	if _, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}); err == nil {
		t.Fatalf("Object was not deleted from S3")
	}
}

// helper functions
func CreateTestFile(uploadFilePath string, t *testing.T) {
	content := []byte("S3 Upload Download Test :3 ^_^")
	if err := os.WriteFile(uploadFilePath, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
}

func FilesAreIdentical(file1, file2 string) bool {
	content1, err := os.ReadFile(file1)
	if err != nil {
		return false
	}
	content2, err := os.ReadFile(file2)
	if err != nil {
		return false
	}
	return bytes.Equal(content1, content2)
}

func CreateRequestBody(bucketName string, objectKey string, filePath string) *bytes.Buffer {
	data := map[string]string{
		"bucketName": bucketName,
		"objectKey":  objectKey,
		"filePath":   filePath,
	}
	body, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(body)
}
