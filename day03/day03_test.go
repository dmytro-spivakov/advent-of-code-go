package day03

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11":                      4361,
		"test_input12":                      6269,
		"test_input13_goddamn_fucking_edge": 4747,
	}

	for inputFile, expectedResult := range cases {
		result := Solution1(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution1() = %v, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution1() = %v, OK\n", result)
		}
	}
}

func TestSolution2(t *testing.T) {
}
