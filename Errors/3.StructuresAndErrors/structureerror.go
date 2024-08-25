package main

import (
	"fmt"
)

type MyError struct {
	DBName string
}

func (e *MyError) Error() string {
	return "Connection to " + e.DBName + " failed"
}

func checkConnection(db string) error {
	return &MyError{DBName: db}
}

func main() {
	var err error

	err = checkConnection("MyDB")
	if err != nil {
		fmt.Println(err)
	}
}
