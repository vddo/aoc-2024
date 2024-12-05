package main

import "testing"

var cachedCsv [][]int

func findByIndex(ind int) []int {
	if len(cachedCsv) == 0 {
		cachedCsv, _ = importCSV2Ints("input.csv")
	}

	return cachedCsv[ind]
}

func TestImportCSV(t *testing.T) {
	type testCase struct {
		name string
		ind  int
		a, b int
	}

	var testPairs []testCase

	testPairs = append(testPairs, testCase{"First Entry", 0, 98415, 86712})
	testPairs = append(testPairs, testCase{"Some where in the middle", 200, 12172, 39516})
	testPairs = append(testPairs, testCase{"Bottom Third", 666, 63810, 76750})
	testPairs = append(testPairs, testCase{"Bottom", 999, 83205, 39489})

	for _, tt := range testPairs {
		t.Run(tt.name, func(t *testing.T) {
			if got := findByIndex(tt.ind); got[0] != tt.a || got[1] != tt.b {
				t.Errorf("At index %d found value pair [%d, %d]; Expected [%d, %d]", tt.ind, got[0], got[1], tt.a, tt.b)
			}
		})
	}
}
