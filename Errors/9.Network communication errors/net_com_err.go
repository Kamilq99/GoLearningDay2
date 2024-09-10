package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// A function that makes an HTTP request and handles errors
func makeHTTPRequest(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Println("Response Body:", string(body))
	return nil
}

func main() {
	// Example URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	err := makeHTTPRequest(url)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Request completed successfully")
	}
}
