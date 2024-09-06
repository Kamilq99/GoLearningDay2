package main

import (
	"errors"
	"fmt"
	"os"
)

// Funkcja próbująca otworzyć plik
func openFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return err
	}
	return nil
}

// Funkcja sprawdzająca, czy błąd jest typu os.ErrNotExist
func checkFile(filename string) error {
	err := openFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		// Jeśli błąd to os.ErrNotExist, wyświetlamy komunikat
		fmt.Printf("Plik %s nie istnieje\n", filename)
		return err
	}
	return err
}

func main() {
	// Próbujemy otworzyć nieistniejący plik
	err := checkFile("nonexistent.txt")
	if err != nil {
		fmt.Println("Błąd:", err)
	}
}
