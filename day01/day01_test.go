package day01

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 142,
		"test_input12": 110,
		"test_input13": 201,
	}

	for inputFile, expectedResult := range cases {
		result := Solution1(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution1() = %v, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution2() = %v, OK\n", result)
		}
	}
}

func TestSolution2(t *testing.T) {
	cases := map[string]int{
		"test_input21": 281,
		"test_input22": 155,
	}

	for inputFile, expectedResult := range cases {
		result := Solution2(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution2() = %v, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution2() = %v, OK\n", result)
		}
	}
}
