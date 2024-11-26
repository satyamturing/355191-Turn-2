package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a POST request
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	r.ParseMultipartForm(32 << 20) // 32MB buffer for file upload
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file to disk
	dest := filepath.Join("uploads", handler.Filename)
	out, err := os.Create(dest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the file content to the destination file
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	func (downloadFileHandler(w http.ResponseWriter, r *http.Request) {
		// Extract the file ID from the URL
		fileID := r.URL.Path[len("/files/"):] // Get file ID from URL
		if fileID == "" {
			http.Error(w, "File ID is required", http.StatusBadRequest)
			return
		}
	
		// Determine the file path
		filePath := filepath.Join("uploads", fileID) // Simulate looking for the file by ID
	
		// Open the file for reading
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		defer file.Close()
	
		// Set headers for file download
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileID)
	
		// Copy the file contents to the response
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Error downloading file", http.StatusInternalServerError)
			return
		}
	}



}
