package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

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

func main() {
	fmt.Println("Start of program door 2")

	// TODO: Import data from file
	data, err := importFile("input.csv")
	if err != nil {
		log.Fatalf("inporting CSV: %v", err)
	}
	fmt.Println(data)

	// TODO: For each row check if state is safe
}
