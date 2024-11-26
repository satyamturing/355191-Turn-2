package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestUploadFileHandler(t *testing.T) {
	// Create a temporary file with sample data
	tempFile, err := ioutil.TempFile("", "testdata")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write sample data to the file
	_, err = tempFile.Write([]byte("Some large data here"))
	if err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}

	// Prepare form data for the POST request
	formData := map[string]string{
		"file": tempFile.Name(),
	}

	// Create the HTTP request to upload the file
	req, err := http.NewRequest("POST", "/files", strings.NewReader(url.FormEncode(formData)))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")

	// Create the recorder to capture the response
	w := httptest.NewRecorder()

	// Set up the router and handler
	mux := http.NewServeMux()
	mux.HandleFunc("/files", uploadFileHandler)

	// Serve the HTTP request and capture the response
	mux.ServeHTTP(w, req)

	// Assertions for the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
	func TestDownloadFileHandler(t *testing.T) {
		// Create a temporary file to simulate an uploaded file
		tempFile, err := ioutil.TempFile("", "uploadedfile")
		if err != nil {
			t.Fatalf("Error creating temp file: %v", err)
		}
		defer os.Remove(tempFile.Name())
	
		// Simulate uploading the file first
		fileID := "1234" // Example file ID
	
		// Create a GET request to download the file
		req, err := http.NewRequest("GET", "/files/"+fileID, nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
	
		// Create the recorder to capture the response
		w := httptest.NewRecorder()
	
		// Set up the router and handler
		mux := http.NewServeMux()
		mux.HandleFunc("/files/", downloadFileHandler)
	
		// Serve the request and capture the response
		mux.ServeHTTP(w, req)
	
		// Assertions for response (e.g., status code, file content)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %d", w.Code)
		}
	}
	





}
