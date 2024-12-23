package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
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

// Take array of arrays with each two elements and split in two arrays with each only one element.
func splitArrayInTwo(array [][]int) ([]int, []int, error) {
	if len(array[0]) != 2 {
		return nil, nil, errors.New("tupel must consist of two elements")
	}
	var array1, array2 []int
	var res [][]int

	for ind := range array {
		array1 = append(array1, array[ind][0])
		array2 = append(array2, array[ind][1])
	}

	res = append(res, array1)
	res = append(res, array2)

	return res[0], res[1], nil
}

func sortSliceOfInts(sliceCoordinates []int) error {
	if len(sliceCoordinates) == 0 {
		return errors.New("slice was empty")
	}

	if slices.IsSorted(sliceCoordinates) {
		return nil
	}

	slices.Sort(sliceCoordinates)

	return nil
}

// Calculate total distance as a sum between the differences of each elements of list 1 and list 2.
func calc(s1, s2 []int) (int, error) {
	if len(s1) != len(s2) {
		return 0, errors.New("arrays are not of the same size")
	}

	sum := 0
	for i := range s1 {
		sum += int(math.Abs(float64(s1[i] - s2[i])))
	}
	return sum, nil
}

func main() {
	fmt.Println("=== Program starts here ===")

	fmt.Println("Read \"input.csv\"")
	csvInput, _ := importCsvToArray("input.csv")
	s1, s2, e := splitArrayInTwo(csvInput)
	if e != nil {
		log.Fatal(e.Error())
	}
	sortSliceOfInts(s1)
	sortSliceOfInts(s2)

	res, e := calc(s1, s2)
	if e != nil {
		log.Fatal(e.Error())
	}

	fmt.Printf("The total distance is %d\n", res)
}
