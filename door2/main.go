package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type LevelData struct {
	Levels     []int
	LineNumber int
}

// Take a CSV file and return a slice of slices with ints.
// The inner slices consist of the data and are of different length.
// Converts elements from string to int
func importFile(in string) ([][]int, error) {
	if in == "" {
		return nil, errors.New("input file name is empty")
	}

	f, e := os.Open(in)
	if e != nil {
		return nil, fmt.Errorf("failed to open file: %w", e)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1

	data := make([][]int, 0, 100)
	trackedLine := 0

	for {
		trackedLine += 1
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading record at line %d: %w", trackedLine, err)
		}

		row := make([]int, 0, 5)
		for i, v := range record {
			vConv, e := strconv.Atoi(v)
			if e != nil {
				return nil, fmt.Errorf("could not convert element %d in line %d: %w", i, trackedLine, e)
			}
			row = append(row, vConv)
		}

		if len(row) == 0 {
			continue
		}
		data = append(data, row)

	}
	return data, nil
}

// Checks first condition: levels (numbers) are all either increasing or decreasing
// Checks second condition: adjacent levels differ max by 3
func checkConditions(row []int) (bool, error) {
	if len(row) == 0 {
		return false, errors.New("row is empty")
	}

	safe := true
	increasing, decreasing, safeDiff := 1, 1, 1
	diff := 0.0

	for i, v := range row {
		if i == len(row)-1 {
			break
		}

		diff = math.Abs(float64(row[i+1] - v))
		if diff > 3 {
			safeDiff = 0
		}

		if v <= row[i+1] {
			increasing = 0
		}

		if v >= row[i+1] {
			decreasing = 0
		}
	}

	if (increasing != 1 && decreasing != 1) || safeDiff != 1 {
		safe = false
	}

	return safe, nil
}

func main() {
	fmt.Println("Start of program door 2")

	data, err := importFile("input.csv")
	if err != nil {
		log.Fatalf("inporting CSV: %v", err)
	}

	countSafe := 0
	for i, row := range data {
		safe, e := checkConditions(row)
		if e != nil {
			log.Fatalf("processing row %d: %v", i, e)
		}

		if safe {
			countSafe += 1
		}
	}

	fmt.Printf("Count of safe reports is \n%d\n", countSafe)
}
