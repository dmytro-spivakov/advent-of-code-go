package day07

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 6440,
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
		"test_input11": 5905,
		"test_input21": 2742,
		"test_input22": 6839,
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

func TestMakeHand(t *testing.T) {
	result := makeHand("AAKKQ", "999", false)

	if result.combinationStrength == 200 && result.bid == 999 && slices.Compare(result.cards, []string{"A", "A", "K", "K", "Q"}) == 0 {
		fmt.Println("makeHand() OK")
	} else {
		t.Fatalf("makeHand() unexpected result %v\n", result)
	}

	result = makeHand("8JJJJ", "999", true)

	if result.combinationStrength == 100000 && result.bid == 999 && slices.Compare(result.cards, []string{"8", "J", "J", "J", "J"}) == 0 {
		fmt.Println("makeHand(joker=true) OK")
	} else {
		t.Fatalf("makeHand(joker=true) unexpected result %v\n", result)
	}
}

func TestCalcCombinationStrength(t *testing.T) {
	cases := map[string]int{
		"AAAAA": 100000, // five of a kind
		"KKJKK": 10000,  // four of a kind
		"JJQQJ": 1100,   // full house
		"AKQKK": 1000,   // three of a kind
		"77AQQ": 200,    // two pairs
		"22AQK": 100,    // one pair
		"AKQJ9": 1,      // highest card
	}

	casesWithJoker := map[string]int{
		"8JJJJ": 100000,
		"7777J": 100000,
		"AABBJ": 1100,
		"7KKJ2": 1000,
	}

	for inputCards, expectedResult := range cases {
		input := strings.Split(inputCards, "")

		result := calcCombinationStrength(input, false)
		if result != expectedResult {
			t.Fatalf("calcCombinationStrength() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("calcCombinationStrength() = %d, OK\n", result)
		}
	}

	for inputCards, expectedResult := range casesWithJoker {
		input := strings.Split(inputCards, "")

		result := calcCombinationStrength(input, true)
		if result != expectedResult {
			t.Fatalf("calcCombinationStrength(joker=true) = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("calcCombinationStrength(joker=true) = %d, OK\n", result)
		}
	}
}
