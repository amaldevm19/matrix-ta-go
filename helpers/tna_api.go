package helpers

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"

	"github.com/amaldevm19/go_matrix_tna/config"
)

func InsertItem(query string) (bool, string, error) {
	// Construct the URL from environment variable and query
	url := config.Config("TNA_URL") + query

	// Read the username and password from environment variables
	tna_username := config.Config("TNA_USERNAME")
	tna_password := config.Config("TNA_PASSWORD")

	// Create the Basic Authorization header value
	auth := base64.StdEncoding.EncodeToString([]byte(tna_username + ":" + tna_password))

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP GET request with the URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, "", err
	}

	// Add the Authorization header to the request
	req.Header.Add("Authorization", "Basic "+auth)

	// Execute the request
	res, err := client.Do(req)

	if err != nil {
		return false, "", nil
	}

	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, "", err
	}

	// Convert the response body to a string
	response_text := string(body)
	if res.StatusCode == http.StatusOK && strings.Contains(response_text, "successful") {
		return true, response_text, nil
	} else {
		return false, response_text, nil
	}

}
