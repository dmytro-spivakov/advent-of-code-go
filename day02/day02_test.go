package day02

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 8,
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

func TestMakeGame(t *testing.T) {
	cases := map[string]game{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   {id: 1, maxRed: 4, maxGreen: 2, maxBlue: 6},
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         {id: 2, maxRed: 1, maxGreen: 3, maxBlue: 4},
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": {id: 3, maxRed: 20, maxGreen: 13, maxBlue: 6},
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": {id: 4, maxRed: 14, maxGreen: 3, maxBlue: 15},
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   {id: 5, maxRed: 6, maxGreen: 3, maxBlue: 2},
	}

	for inputString, expectedResult := range cases {
		newGame := makeGame(inputString)

		if newGame != expectedResult {
			t.Fatalf("makeGame() unexpected results for game %v\n", expectedResult.id)
		} else {
			fmt.Printf("makeGame(), game %v OK\n", newGame.id)
		}
	}
}

func TestSolution2(t *testing.T) {
}
