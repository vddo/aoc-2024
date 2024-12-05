package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Program starts here")

	// Open the file
	file, error := os.Open("input.csv")
	if error != nil {
		log.Fatal(error)
	}

	file.Close()

	// Create a new CSV Reader
	reader := csv.NewReader(file)

	// Read the records
}
