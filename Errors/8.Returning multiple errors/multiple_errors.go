package main

import (
	"fmt"
	"strings"
)

// Custom error type to hold multiple errors
type MultiError struct {
	Errors []error
}

// Implementing the Error interface for MultiError
func (m *MultiError) Error() string {
	var errStrings []string
	for _, err := range m.Errors {
		errStrings = append(errStrings, err.Error())
	}
	return strings.Join(errStrings, "; ")
}

// A function that returns multiple errors
func processValues(values []int) error {
	var multiErr MultiError

	for _, value := range values {
		if value < 0 {
			multiErr.Errors = append(multiErr.Errors, fmt.Errorf("negative value: %d", value))
		}
	}

	// Return nil if no errors, otherwise return MultiError
	if len(multiErr.Errors) == 0 {
		return nil
	}
	return &multiErr
}

func main() {
	// Example input with multiple errors
	values := []int{10, -1, 5, -3, 7}
	err := processValues(values)
	if err != nil {
		fmt.Println("Errors occurred:", err)
	} else {
		fmt.Println("Processing completed successfully")
	}
}
