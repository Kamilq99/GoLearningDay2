package main

import (
	"errors"
	"fmt"
)

func openFile(name string) (string, error) {
	if name == "error.txt" {
		return "", errors.New("File does not exist")
	}
	return "File opened successfully", nil
}

func main() {

	//file, err := openFile("Text.docx")
	file, err := openFile("error.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
}
