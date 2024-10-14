package day09

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 114,
	}

	for inputFile, expectedResult := range cases {
		result := Solution1(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution1() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution1() = %d, OK\n", result)
		}
	}
}

func TestSolution2(t *testing.T) {
	cases := map[string]int{
		"test_input11": 2,
	}

	for inputFile, expectedResult := range cases {
		result := Solution2(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution2() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution2() = %d, OK\n", result)
		}
	}
}
