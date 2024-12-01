package day10

import (
	"fmt"
	"slices"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 4,
		"test_input12": 8,
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
		"test_input21": 4,
		"test_input22": 8,
		"test_input23": 10,
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

func TestFindAdjacent(t *testing.T) {
	inputMatrix, _ := readMatrix("test_input11")
	inputs := [][2]int{{1, 3}, {3, 1}}
	expectedResults := [][][2]int{
		{{1, 2}, {2, 3}},
		{{2, 1}, {3, 2}},
	}

	for i, input := range inputs {
		results := findAdjacent(&inputMatrix, input)
		expected := expectedResults[i]

		if slices.Equal(expected, results) {
			fmt.Println("findAdjacent() OK")
		} else {
			t.Fatalf("findAdjacent() = %v, expecting %v\n", results, expected)
		}
	}
}
