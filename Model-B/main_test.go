package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestProcessDataHandler(t *testing.T) {
	// Prepare test data (large dataset)
	testData := make([]int, 10000)
	for i := range testData {
		testData[i] = i
	}
	testJSON, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Error marshaling test data: %v", err)
	}

	// Define pagination parameters
	page := 2
	pageSize := 10
	params := fmt.Sprintf("?page=%d&pageSize=%d", page, pageSize)

	// Create a request to the API endpoint
	req := httptest.NewRequest("POST", fmt.Sprintf("/processData%s", params), bytes.NewReader(testJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(processData)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", status)
	}

	// Decode the response body
	var response DataResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Assert the response
	expectedData := testData[(page-1)*pageSize : page*pageSize]
	if !reflect.DeepEqual(response.Data, expectedData) {
		t.Errorf("Expected data %v, got %v", expectedData, response.Data)
	}

	if response.Page != page {
		t.Errorf("Expected page %d, got %d", page, response.Page)
	}

	if response.PageSize != pageSize {
		t.Errorf("Expected page size %d, got %d", pageSize, response.PageSize)
	}

	if response.Total != len(testData) {
		t.Errorf("Expected total %d, got %d", len(testData), response.Total)
	}
}
