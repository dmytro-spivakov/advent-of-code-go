package day15

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 1320,
	}

	for input, expectedResult := range cases {
		result := Solution1(input)

		if result == expectedResult {
			fmt.Printf("Solution1()=%d, OK\n", result)
		} else {
			t.Fatalf("Solution1()=%d, expecting %d, FAIL\n", result, expectedResult)
		}
	}
}

func TestSegmentHash(t *testing.T) {
	cases := map[string]int{
		"rn=1": 30,
		"cm-":  253,
		"qp=3": 97,
		"cm=2": 47,
		"qp-":  14,
		"pc=4": 180,
		"ot=9": 9,
		"ab=5": 197,
		"pc-":  48,
		"pc=6": 214,
		"ot=7": 231,
	}

	for input, expectedResult := range cases {
		result := segmentHash(input)

		if result == expectedResult {
			fmt.Printf("segmentHash()=%d, OK\n", result)
		} else {
			t.Fatalf("segmentHash()=%d, expecting %d, FAIL\n", result, expectedResult)
		}
	}
}
