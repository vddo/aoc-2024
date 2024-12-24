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

const (
	expectedColumns = 2
	initialCap      = 100
)

func importCsvToArray(input string) ([][]int, error) {
	if input == "" {
		return nil, errors.New("empty input filename")
	}

	file, error := os.Open(input)
	if error != nil {
		return nil, fmt.Errorf("opening file: %w", error)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.ReuseRecord = true
	// Performance: skipp early re-allocations
	csvInput := make([][]int, 0, initialCap)
	trackRow := 0

	for {
		trackRow += 1
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("reading CSV rows %d: %w", trackRow, err)
		}

		if len(row) != 1 {
			return nil, fmt.Errorf("reading CSV rows %d: found %d columns; expect 1", trackRow, len(row))
		}

		fields := strings.Fields(row[0])
		if len(fields) != expectedColumns {
			return nil, fmt.Errorf("extracted CSV rows %d: found %d fields; expect %d", trackRow, len(fields), expectedColumns)
		}
		rowValues := make([]int, expectedColumns)

		for ind, val := range fields {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("converting %s in row %d to int: %w", val, trackRow, err)
			}

			rowValues[ind] = valInt
		}

		csvInput = append(csvInput, rowValues)
	}

	fmt.Println("CSV file read successfully.")
	return csvInput, nil
}

// Take array of arrays with each two elements and split in two arrays with each only one element.
func splitArrayInTwo(array [][]int) ([]int, []int, error) {
	if len(array[0]) != expectedColumns {
		return nil, nil, errors.New("tupel must consist of two elements")
	}
	array1 := make([]int, len(array))
	array2 := make([]int, len(array))

	for i, pair := range array {
		array1[i] = pair[0]
		array2[i] = pair[1]
	}

	return array1, array2, nil
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
