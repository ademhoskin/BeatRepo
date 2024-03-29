package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// test config

var (
	bucketName        = "cppbeatproj"
	uploadObjectKey   = "test-upload.txt"
	uploadFilePath    = "test-upload.txt"
	downloadObjectKey = "1347872.jpg"
	downloadFilePath  = "1347872.jpg"
)

func TestThisConfig(t *testing.T) {
	if bucketName == "" || uploadObjectKey == "" || uploadFilePath == "" || downloadObjectKey == "" || downloadFilePath == "" {
		t.Fatalf("Test config is not set")
	}
}

func TestS3UploadDownloadIntegration(t *testing.T) {
	CreateTestFile(uploadFilePath, t)
	defer os.Remove(uploadFilePath)

	// Upload the file to S3
	if err := S3UploadFile(bucketName, uploadObjectKey, uploadFilePath); err != nil {
		t.Fatalf("Failed to upload test file to S3: %v", err)
	}

	//defer the deletion of the test object from S3
	defer func() {
		if err := S3DeleteFile(bucketName, uploadObjectKey); err != nil {
			t.Fatalf("Failed to delete test object from S3: %v", err)
		}
	}()

	// Download the file from S3
	if err := S3DownloadFile(bucketName, uploadObjectKey, downloadFilePath); err != nil {
		t.Fatalf("Failed to download test file from S3: %v", err)
	}
	defer os.Remove(downloadFilePath)

}

func TestS3UploadHandler(t *testing.T) {
	CreateTestFile(uploadFilePath, t)
	defer os.Remove(uploadFilePath)
	defer func() {
		if err := S3DeleteFile(bucketName, uploadObjectKey); err != nil {
			t.Fatalf("Failed to delete test object from S3: %v", err)
		}
	}()

	body := CreateRequestBody(bucketName, uploadObjectKey, uploadFilePath)
	if body == nil {
		t.Fatalf("Failed to create request body")
	}

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestS3DownloadHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/download?bucketName="+bucketName+"&objectKey="+downloadObjectKey+"&filePath="+downloadFilePath, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DownloadHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	defer os.Remove(downloadFilePath)
}

func TestS3DeleteHandler(t *testing.T) {
	CreateTestFile(uploadFilePath, t)
	defer os.Remove(uploadFilePath)

	if err := S3UploadFile(bucketName, uploadObjectKey, uploadFilePath); err != nil {
		t.Fatalf("Failed to upload test file(upload) to S3: %v", err)
	}

	deleteURL := "/delete?bucketName=" + bucketName + "&objectKey=" + uploadObjectKey
	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if err := S3DownloadFile(bucketName, uploadObjectKey, uploadFilePath); err == nil {
		t.Fatalf("Failed to delete test object from S3")
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
