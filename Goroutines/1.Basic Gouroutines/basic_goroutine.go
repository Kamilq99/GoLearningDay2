package main

import (
	"fmt"
	"time"
)

func runTimer1() {
	array := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(array); i++ {
		fmt.Println("Goroutine 1:", array[i])
		time.Sleep(time.Second)
	}
}

func runTimer2() {
	array := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(array); i++ {
		fmt.Println("Goroutine 2:", array[i])
		time.Sleep(time.Second)
	}
}

func main() {
	go runTimer1()
	go runTimer2()

	time.Sleep(time.Second * 6)
	fmt.Println("Time is up!")
}
