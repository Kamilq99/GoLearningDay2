package main

import (
	"errors"
	"fmt"
)

// Define a custom error for demonstration
var ErrNotFound = errors.New("resource not found")

// A function that returns a wrapped error
func findResource(id string) error {
	// Simulate that the resource is not found
	if id == "" {
		return fmt.Errorf("failed to find resource: %w", ErrNotFound)
	}
	return nil
}

// A function that checks if the error is of the same type (ErrNotFound)
func isNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func main() {
	// Test case where the resource is not found (empty id)
	err := findResource("")
	if err != nil {
		fmt.Println("Error:", err)

		// Check if the error is ErrNotFound
		if isNotFoundError(err) {
			fmt.Println("The error is of type ErrNotFound")
		} else {
			fmt.Println("The error is of a different type")
		}
	} else {
		fmt.Println("Resource found successfully!")
	}
}
