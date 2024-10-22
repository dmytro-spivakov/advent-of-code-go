package day12

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
)

func TestSolutio1(t *testing.T) {
	cases := map[string]int{
		"test_input12": 6,
		"test_input11": 21,
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

func TestSolutio1Alt(t *testing.T) {
	cases := map[string]int{
		"test_input13": 1,
		"test_input12": 6,
		"test_input11": 21,
	}

	for input, expectedResult := range cases {
		result := Solution1Alt(input)

		if result == expectedResult {
			fmt.Printf("Solution1Alt()=%d, OK\n", result)
		} else {
			t.Fatalf("Solution1Alt()=%d, expecting %d, FAIL\n", result, expectedResult)
		}
	}
}

func TestGetAllPossibleStrings(t *testing.T) {
	input := "?..?##.?"
	expectedResult := []string{
		"#..###.#",
		"#..###..",
		"#...##.#",
		"#...##..",
		"...###.#",
		"...###..",
		"....##.#",
		"....##..",
	}
	chunkVars := map[int][]string{
		0: {".", "#"},
		3: {".", "#"},
		7: {".", "#"},
	}

	result := getAllPossibleStrings(input, chunkVars, []int{0, 3, 7})
	slices.Sort(expectedResult)
	slices.Sort(result)
	if slices.Equal(result, expectedResult) {
		fmt.Printf("getAllPossibleStrings()=%v; OK\n", result)
	} else {
		t.Fatalf("getAllPossibleStrings()=%v, expecting %v; FAIL\n", result, expectedResult)
	}
}

func TestGetAllPossibleCombinations(t *testing.T) {
	cases := map[int][]string{
		1: {".", "#"},
		2: {
			"..",
			"#.",
			".#",
			"##",
		},
	}

	for input, expectedResult := range cases {
		result := getPossibleCombinations(input)
		slices.Sort(result)
		slices.Sort(expectedResult)

		if slices.Equal(result, expectedResult) {
			fmt.Printf("getPossibleCombinations()=%v, OK\n", result)
		} else {
			t.Fatalf("getPossibleCombinations()=%v, expecting %v, FAIL\n", result, expectedResult)
		}
	}
}

func TestGetAllUnsolvedChunkRanges(t *testing.T) {
	cases := map[string]map[int]int{
		"???.###":        {0: 3},
		".??..??...?##.": {1: 2, 5: 2, 10: 1},
		"?#?#?":          {0: 1, 2: 1, 4: 1},
		"?###???????":    {0: 1, 4: 7},
	}

	for input, expectedResult := range cases {
		result := getUnsolvedChunkRanges(strings.Split(input, ""))

		if maps.Equal(result, expectedResult) {
			fmt.Printf("getUnsolvedChunkRanges()=%v; OK\n", result)
		} else {
			t.Fatalf("getUnsolvedChunkRanges()=%v, expecting %v; FAIL\n", result, expectedResult)
		}
	}
}
