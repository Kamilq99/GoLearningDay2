package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
		wg.Done()
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(counter)
}