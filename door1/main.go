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

func importCsvToArray(input string) ([][]int, error) {
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

	for {
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

	fmt.Println("CSV file read successfully.")
	return csvInput, nil
}

func splitArrayInTwo(array [][]int) ([][]int, error) {
	// Take array of arrays with each two elements and split in two arrays with each only one element.
	var array1, array2 []int
	var res [][]int

	for ind := range array {
		array1 = append(array1, array[ind][0])
		array2 = append(array2, array[ind][1])
	}

	res = append(res, array1)
	res = append(res, array2)

	// TODO: Check input array
	return res, nil
}

func main() {
	fmt.Println("=== Program starts here ===")

	fmt.Println("Read \"input.csv\"")
	csvInput, _ := importCsvToArray("input.csv")
	splitArray, _ := splitArrayInTwo(csvInput)
	fmt.Println(splitArray[0])
}
