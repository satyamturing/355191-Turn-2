package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// DataRequest represents the request body structure
type DataRequest struct {
	Data []int `json:"data"`
}

// DataResponse represents the response body structure
type DataResponse struct {
	Data     []int `json:"data"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Total    int   `json:"total"`
}

// processData handles the POST request to /processData
func processData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body
	var req DataRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Extract pagination parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Calculate pagination indices
	total := len(req.Data)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	// Prepare the response data
	res := DataResponse{
		Data:     req.Data[start:end],
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/processData", processData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
