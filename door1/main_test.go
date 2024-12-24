package main

import (
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

// func TestSorting(t *testing.T) {
// 	// Test setup
// 	csv, err := importCsvToArray("input.csv")
// 	if err != nil {
// 		t.Fatalf("Failed to import CSV: %v", err)
// 	}
//
// 	// Store original data for comparison
// 	originalData := make([][]int, len(csv))
// 	copy(originalData, csv)
//
// 	// Test array splitting
// 	array1, array2, err := splitArrayInTwo(csv)
// 	if err != nil {
// 		t.Fatalf("Failed to split arrays: %v", err)
// 	}
//
// 	// Verify split was correct
// 	if len(array1)+len(array2) != len(csv) {
// 		t.Errorf("Split arrays length mismatch: got %d + %d = %d, want total of %d",
// 			len(array1), len(array2), len(array1)+len(array2), len(csv))
// 	}
//
// 	// Test sorting of first array
// 	if err := sortSliceOfInts(array1); err != nil {
// 		t.Fatalf("Failed to sort first array: %v", err)
// 	}
//
// 	// Test sorting of second array
// 	if err := sortSliceOfInts(array2); err != nil {
// 		t.Fatalf("Failed to sort second array: %v", err)
// 	}
//
// 	// Verify sorting results
// 	t.Run("Verify Array1 Sorting", func(t *testing.T) {
// 		if !slices.IsSorted(array1) {
// 			t.Error("First array is not sorted")
// 			if len(array1) > 1 {
// 				// Print a snippet of the unsorted sequence
// 				for i := 0; i < len(array1)-1; i++ {
// 					if array1[i] > array1[i+1] {
// 						t.Errorf("Unsorted sequence found at index %d: %v > %v",
// 							i, array1[i], array1[i+1])
// 						break
// 					}
// 				}
// 			}
// 		}
// 	})
//
// 	t.Run("Verify Array2 Sorting", func(t *testing.T) {
// 		if !slices.IsSorted(array2) {
// 			t.Error("Second array is not sorted")
// 			if len(array2) > 1 {
// 				for i := 0; i < len(array2)-1; i++ {
// 					if array2[i] > array2[i+1] {
// 						t.Errorf("Unsorted sequence found at index %d: %v > %v",
// 							i, array2[i], array2[i+1])
// 						break
// 					}
// 				}
// 			}
// 		}
// 	})
//
// 	// Verify no data was lost during sorting
// 	t.Run("Verify Data Integrity", func(t *testing.T) {
// 		combinedLen := len(array1) + len(array2)
// 		if combinedLen != len(originalData) {
// 			t.Errorf("Data loss detected: original length %d, current length %d",
// 				len(originalData), combinedLen)
// 		}
//
// 		// Optional: verify all original values are present in sorted arrays
// 		if !verifyDataIntegrity(originalData, array1, array2) {
// 			t.Error("Sorted arrays don't contain all original values")
// 		}
// 	})
// }
//
// // Helper function to verify data integrity
// func verifyDataIntegrity(original [][]int, arr1, arr2 [][]int) bool {
// 	// Create maps to count occurrences of values
// 	originalCount := make(map[string]int)
// 	sortedCount := make(map[string]int)
//
// 	// Count original values
// 	for _, v := range original {
// 		key := fmt.Sprintf("%v", v)
// 		originalCount[key]++
// 	}
//
// 	// Count values in sorted arrays
// 	for _, v := range arr1 {
// 		key := fmt.Sprintf("%v", v)
// 		sortedCount[key]++
// 	}
// 	for _, v := range arr2 {
// 		key := fmt.Sprintf("%v", v)
// 		sortedCount[key]++
// 	}
//
// 	// Compare counts
// 	return reflect.DeepEqual(originalCount, sortedCount)
// }
