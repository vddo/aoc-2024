package main

import (
	"door4/importdata"
	"door4/solver"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Program door4 starting...")

	data, err := importdata.Import("input-small")
	if err != nil {
		log.Fatal(err)
	}

	s := solver.NewSolver(data, solver.KEYWORD, len(*data), len((*data)[0]))
	s_err := s.Solve()
	if err != nil {
		log.Fatalf("Not able to solve this input: %v", s_err)
	}

	s.Render()
}
