package day04

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 13,
	}

	for inputFile, expectedResult := range cases {
		result := Solution1(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution1() = %v, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution1() = %v, OK\n", result)
		}
	}
}

func TestParseCard(t *testing.T) {
	cases := map[string]card{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53": {
			id:             1,
			winningNumbers: map[int]bool{41: true, 48: true, 83: true, 86: true, 17: true},
			yourNumbers:    map[int]bool{83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true},
		},
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1": {
			id:             3,
			winningNumbers: map[int]bool{1: true, 21: true, 53: true, 59: true, 44: true},
			yourNumbers:    map[int]bool{69: true, 82: true, 63: true, 72: true, 16: true, 21: true, 14: true, 1: true},
		},
	}

	for inputString, expectedCard := range cases {
		result := parseCard(inputString)
		if result.id != expectedCard.id || !maps.Equal(result.winningNumbers, expectedCard.winningNumbers) || !maps.Equal(result.yourNumbers, expectedCard.yourNumbers) {
			t.Fatalf("parseCard() = %v, expecting %v\n", result, expectedCard)
		} else {
			fmt.Printf("parseCard() = %v, OK\n", result)
		}
	}
}

func TestCardPoints(t *testing.T) {
	cards := []card{
		{
			id:             1,
			winningNumbers: map[int]bool{41: true, 48: true, 83: true, 86: true, 17: true},
			yourNumbers:    map[int]bool{48: true, 83: true, 86: true, 17: true, 6: true, 31: true, 9: true, 53: true},
		},
		{
			id:             2,
			winningNumbers: map[int]bool{1: true, 21: true, 53: true, 59: true, 44: true},
			yourNumbers:    map[int]bool{69: true, 82: true, 63: true, 72: true, 16: true, 21: true, 14: true, 3333: true},
		},
		{
			id:             3,
			winningNumbers: map[int]bool{1: true, 21: true, 53: true, 59: true, 44: true},
			yourNumbers:    map[int]bool{69: true, 82: true, 63: true, 72: true, 16: true, 21: true, 14: true, 1: true},
		},
		{
			id:             4,
			winningNumbers: map[int]bool{1: true, 21: true, 53: true, 59: true, 44: true},
			yourNumbers:    map[int]bool{69: true, 82: true, 63: true, 72: true, 16: true, 23: true, 14: true, 2: true},
		},
	}
	cases := map[int]int{
		1: 8,
		2: 1,
		3: 2,
		4: 0,
	}

	for cardId, expectedResult := range cases {
		cardIndex := slices.IndexFunc(cards, func(c card) bool {
			return c.id == cardId
		})
		card := cards[cardIndex]

		result := card.points()
		if result != expectedResult {
			t.Fatalf("card.points() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("card.points() = %d, OK\n", result)
		}
	}
}
