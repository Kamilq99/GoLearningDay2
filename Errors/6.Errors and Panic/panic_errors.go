package main

import (
	"errors"
	"fmt"
)

// Panic if attempts exceed maxAttempts
func doTask(attempts, maxAttempts int) {
	if attempts > maxAttempts {
		panic(fmt.Sprintf("Exceeded maximum attempts: %d", maxAttempts))
	}
	fmt.Println("Task completed successfully on attempt", attempts)
}

// Retry logic with panic and recover
func retryTask(maxAttempts int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Recover from panic and return the error
			err = errors.New(fmt.Sprintf("Task failed: %v", r))
		}
	}()

	for attempts := 1; attempts <= maxAttempts+1; attempts++ {
		doTask(attempts, maxAttempts) // Call the function that may panic
	}

	return nil
}

func main() {
	// Set maximum attempts
	maxAttempts := 3

	// Try to perform the task and handle any panic
	err := retryTask(maxAttempts)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Task completed successfully")
	}
}
