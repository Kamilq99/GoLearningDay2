package main

import (
	"errors"
	"fmt"
)

func divideNumbers(a float64, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return a / b, nil
}

func main() {

	var a, b float64

	fmt.Printf("Get a first number: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Get a second number: ")
	fmt.Scanf("%f", &b)

	result, err := divideNumbers(a, b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
