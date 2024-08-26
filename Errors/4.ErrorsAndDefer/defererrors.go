package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("failed to close file: %v\n", cerr)
		}
	}()

	fmt.Println("File opened successfully")

	return nil
}

func main() {
	if err := readFile("example.txt"); err != nil {
		fmt.Println("Error:", err)
	}
}
