package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func importCSV2Ints(input string) ([][]int, error) {
	// Open the file
	file, error := os.Open(input)
	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()

	// Create a new CSV Reader
	reader := csv.NewReader(file)
	reader.ReuseRecord = true

	// Read the records
	var csvInput [][]int
	hackCount := 0

	for {
		hackCount += 1
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		row = strings.Fields(row[0])

		rowInts := make([]int, 2)

		for ind, val := range row {

			valInt, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}

			rowInts[ind] = valInt
		}

		csvInput = append(csvInput, rowInts)
	}

	return csvInput, nil
}

func main() {
	fmt.Println("=== Program starts here ===")

	csvInput, _ := importCSV2Ints("input.csv")
	fmt.Println(csvInput[0])
}
