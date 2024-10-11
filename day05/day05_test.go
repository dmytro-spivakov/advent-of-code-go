package day05

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 3,
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

func TestRangeFindDest(t *testing.T) {
	rangeStr := Range{srcStart: 10, destStart: 20, length: 10}
	cases := map[int]int{
		10: 20,
		20: 30,
		15: 25,
		9:  -1,
		31: -1,
	}

	for srcStart, expectedDest := range cases {
		result := rangeStr.findDest(srcStart)
		if result != expectedDest {
			t.Fatalf("Range.findDest() = %d, expecting %d\n", result, expectedDest)
		} else {
			fmt.Printf("Range.findDest() = %d, OK\n", result)
		}
	}
}
