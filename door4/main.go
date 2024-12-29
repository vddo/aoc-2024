package main

import (
	"door4/importdata"
	"door4/solver"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Program door4 starting...")

	data, err := importdata.Import("input")
	if err != nil {
		log.Fatal(err)
	}

	count, _ := solver.Solve(data)
	fmt.Println(count)
}
