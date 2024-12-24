package main

import (
	"slices"
	"testing"
)

func TestImportCSV(t *testing.T) {
	type testCase struct {
		name string
		ind  int
		a, b int
	}

	csv, e := importCsvToArray("input.csv")
	if e != nil {
		t.Errorf("Failed to inport from input file.")
		t.FailNow()
	}

	testCases := []testCase{
		{"First Entry", 0, 98415, 86712},
		{"Somewhere in the middle", 200, 12172, 39516},
		{"Bottom Third", 666, 63810, 76750},
		{"Bottom", 999, 83205, 39489},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := csv[tt.ind]; got[0] != tt.a || got[1] != tt.b {
				t.Errorf("At index %d found value pair [%d, %d]; Expected [%d, %d]", tt.ind, got[0], got[1], tt.a, tt.b)
			}
		})
	}
}

func TestSorting(t *testing.T) {
	csv, e := importCsvToArray("input.csv")
	if e != nil {
		t.Errorf("Failed to inport from input file.")
		t.FailNow()
	}

	a1, a2, e := splitArrayInTwo(csv)
	if e != nil {
		t.Errorf("Failed to split arrays.")
		t.FailNow()
	}

	e = sortSliceOfInts(a1)
	if e != nil {
		t.Errorf("Failed to sort arrays.")
		t.FailNow()
	}

	e = sortSliceOfInts(a2)
	if e != nil {
		t.Errorf("Failed to sort arrays.")
		t.FailNow()
	}

	if slices.IsSorted(a1) == false || slices.IsSorted(a2) == false {
		t.Error("arrays are not sorted")
	}
}
