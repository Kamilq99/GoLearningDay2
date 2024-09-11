package main

import (
	"fmt"
)

func addNumbers(chanel chan int) {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	chanel <- sum
}

func main() {
	chanel := make(chan int)

	go addNumbers(chanel)

	result := <-chanel

	fmt.Println("Result It is:", result)
}
