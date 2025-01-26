package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	rowRuleData string
	rowLogData  string
)

func main() {
	fmt.Println("Solving door 5...")

	// Buffer for rules and logs
	bufRules := make([]rowRuleData, 0, 100)
	bufLogs := make([]rowLogData, 0, 20)

	f, err := os.Open("data/small.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	flagScanRules := true
	for scanner.Scan() {

		row := scanner.Text()

		if len(row) == 0 {
			flagScanRules = !flagScanRules
			continue
		}

		if flagScanRules {
			bufRules = append(bufRules, rowRuleData(row))
		} else {
			bufLogs = append(bufLogs, rowLogData(row))
		}
	}

	fmt.Printf("Rows of rules: %d\n"+"Rows of logs: %d\n", len(bufRules), len(bufLogs))
}
