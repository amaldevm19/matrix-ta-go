package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/amaldevm19/go_matrix_tna/handler"

	"testing"
)

type BranchPayload struct {
	BranchName string `json:"branch_name"`
	BranchCode string `json:"branch_code"`
}

func TestAddNewBranch(t *testing.T) {

	// Define test cases
	tests := []struct {
		name                     string
		branchName               string
		branchCode               string
		expectedStatus           string
		expectedError            string
		expectedResponseContains string
	}{
		{
			name:                     "successful request",
			branchName:               "branch123",
			branchCode:               "BRC123",
			expectedStatus:           "ok",
			expectedResponseContains: "ok",
		},
		{
			name:                     "non-successful request",
			branchName:               "branch123",
			branchCode:               "BRC1231212121",
			expectedStatus:           "failed",
			expectedResponseContains: "failed",
		},
		{
			name:                     "non-successful request with no branchName and branchCode",
			expectedStatus:           "failed",
			expectedResponseContains: "Values cannot be empty",
		},
		{
			name:                     "non-successful request without branch_name",
			branchName:               "branch123",
			expectedStatus:           "failed",
			expectedResponseContains: "Values cannot be empty",
		},
		{
			name:                     "non-successful request without branch_code",
			branchCode:               "BRC123",
			expectedStatus:           "failed",
			expectedResponseContains: "Values cannot be empty",
		},
		{
			name:                     "non-successful request with empty values",
			branchName:               "",
			branchCode:               "",
			expectedStatus:           "failed",
			expectedResponseContains: "Values cannot be empty",
		},
		{
			name:                     "non-successful request with empty branch_name",
			branchName:               "",
			branchCode:               "BRC123",
			expectedStatus:           "failed",
			expectedResponseContains: "Values cannot be empty",
		},
		{
			name:                     "non-successful request with empty branch_code",
			branchName:               "branch123",
			branchCode:               "",
			expectedStatus:           "failed",
			expectedResponseContains: "Values cannot be empty",
		},
		{
			name:                     "non-successful request with duplicate values",
			branchName:               "branch123",
			branchCode:               "BRC123",
			expectedStatus:           "failed",
			expectedResponseContains: "Code already exists",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a JSON payload
			payload := BranchPayload{
				BranchName: tc.branchName,
				BranchCode: tc.branchCode,
			}
			jsonPayload, err := json.Marshal(payload)

			if err != nil {
				t.Fatalf("error marshalling payload: %v", err)
			}

			// Create a new HTTP client
			client := &http.Client{}

			// Create a POST request
			req, err := http.NewRequest("POST", "http://127.0.0.1:8000/api/branch", bytes.NewBuffer(jsonPayload))
			if err != nil {
				t.Fatalf("error creating request: %v", err)
			}

			// Add headers
			req.Header.Set("Content-Type", "application/json")

			// Execute the request
			res, err := client.Do(req)

			if err != nil {
				t.Fatalf("error in POST[http://127.0.0.1:8000/branch] : %v", err)
			}

			defer res.Body.Close()

			// Read the response body
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("error reading response body: %v", err)
			}

			// Unmarshal the response body into a Response struct
			var response handler.Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Fatalf("error unmarshalling response body: %v", err)
			}

			// Check if the response fields match the expected values
			if response.Status != tc.expectedStatus {
				t.Errorf("expected status '%s', got '%s'", tc.expectedStatus, response.Status)
			}

			// Check if the response body contains the expected substring
			if !bytes.Contains(body, []byte(tc.expectedResponseContains)) {
				t.Errorf("expected response to contain '%s', got '%s'", tc.expectedResponseContains, body)
			}

		})
	}
}
