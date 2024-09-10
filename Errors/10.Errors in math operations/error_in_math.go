package main

import (
	"errors"
	"fmt"
	"math"
)

// Function that returns an error if the operation is not possible
func sqrtWithCheck(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot take the square root of a negative number")
	}
	return math.Sqrt(a), nil
}

// Another operation (division) that returns an error
func divideWithCheck(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Example for square root
	num := -9.0
	result, err := sqrtWithCheck(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Square root of %.2f is %.2f\n", num, result)
	}

	// Example for division
	a, b := 10.0, 0.0
	divResult, err := divideWithCheck(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.2f divided by %.2f is %.2f\n", a, b, divResult)
	}
}
