package day06

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 288,
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
		"test_input11": 71503,
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

func TestRaceCalcNewRecordOptions(t *testing.T) {
	// {race time, record distace, expected number of new record options}
	cases := [][]int{
		[]int{7, 9, 4},
		[]int{15, 40, 8},
		[]int{30, 200, 9},
	}

	for _, testCase := range cases {
		race := Race{time: testCase[0], distance: testCase[1]}
		expectedResult := testCase[2]

		result := race.calcNewRecordOptions()
		if result != expectedResult {
			t.Fatalf("Race.calcNewRecordOptions() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Race.calcNewRecordOptions() = %d, OK\n", result)
		}
	}
}

func TestParseScatteredInt(t *testing.T) {
	cases := map[string]int{
		"7  15   30": 71530,
		"9  40  200": 940200,
	}

	for input, expectedResult := range cases {
		result := parseScatteredInt(input)
		if result != expectedResult {
			t.Fatalf("parseScatteredInt() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("parseScatteredInt() = %d, OK\n", result)
		}
	}
}
