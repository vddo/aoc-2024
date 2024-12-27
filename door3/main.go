package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Return multiplication x by y
func mul(x, y int) int {
	return x * y
}

type lines struct {
	data       string
	lineNumber int
}

type factors struct {
	multiplier   int
	multiplicant int
}

func importData(input string) ([]lines, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	fmt.Println("Start scanning file ...")
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	data := make([]lines, 0, 10)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("scanning file: %w", err)
		}

		line := lines{}
		line.data = scanner.Text()

		if len(line.data) == 0 {
			continue
		}

		lineNumber++
		line.lineNumber = lineNumber
		data = append(data, line)
	}

	fmt.Println("Imported file successfully")

	return data, nil
}

func parseMul(op string) (*factors, error) {
	re := regexp.MustCompile(`\d+`)
	parsed := re.FindAllString(op, -1)
	if len(parsed) != 2 {
		return nil, errors.New("not exactely two numbers")
	}

	m1, err := strconv.Atoi(parsed[0])
	if err != nil {
		return nil, fmt.Errorf("converting string to integer: %w", err)
	}
	m2, err := strconv.Atoi(parsed[1])
	if err != nil {
		return nil, fmt.Errorf("converting string to integer: %w", err)
	}

	return &factors{m1, m2}, nil
}

func parser(input []lines) ([]*factors, error) {
	collectionFactors := make([]*factors, 0, 20)
	collectionParsed := make([]string, 0, 20)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	for i := 0; i < len(input); i++ {
		parsed := re.FindAllString(input[i].data, -1)
		if parsed != nil {
			collectionParsed = append(collectionParsed, parsed...)
		}
	}

	for _, operation := range collectionParsed {
		parsedFactors, err := parseMul(operation)
		if err != nil {
			return nil, fmt.Errorf("parsing a single mul(x,y): %w", err)
		}

		collectionFactors = append(collectionFactors, parsedFactors)
	}

	return collectionFactors, nil
}

// Calculates all mul-operations and sums up all results
func calc(data []*factors) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("empty data provided")
	}

	sum := 0
	for _, mulOperation := range data {
		sum += mulOperation.multiplicant * mulOperation.multiplier
	}

	return sum, nil
}

func main() {
	fmt.Println("Program door 3 started")
	data, err := importData("input")
	if err != nil {
		log.Fatalf("imporing file: %v", err)
	}

	parsedData, err := parser(data)
	if err != nil {
		log.Fatalf("parsing data: %v", err)
	}

	result, _ := calc(parsedData)
	fmt.Printf("The sum of all mul operations is\n%d\n", result)
}
