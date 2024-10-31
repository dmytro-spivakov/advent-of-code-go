package day14

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 136,
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

func TestSolution2(t *testing.T) {
	cases := map[string]int{
		"test_input11": 64,
	}

	for input, expectedResult := range cases {
		result := Solution2(input)

		if result == expectedResult {
			fmt.Printf("Solution2()=%d, OK\n", result)
		} else {
			t.Fatalf("Solution2()=%d, expecting %d, FAIL\n", result, expectedResult)
		}
	}
}

func TestTranspose(t *testing.T) {
}
