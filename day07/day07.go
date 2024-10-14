package day07

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards               []string
	bid                 int
	combinationStrength int
}

func makeHand(cards string, bidStr string, withJokers bool) Hand {
	h := Hand{cards: strings.Split(cards, "")}
	bid, err := strconv.ParseInt(bidStr, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse int %v\n", bidStr)
	}
	h.bid = int(bid)
	h.combinationStrength = calcCombinationStrength(h.cards, withJokers)

	return h
}

func calcCombinationStrength(cards []string, withJokers bool) int {
	cardsCount := make(map[string]int)

	for _, card := range cards {
		cardsCount[card] += 1
	}

	if withJokers && cardsCount["J"] > 0 && cardsCount["J"] < 5 {
		maxCount := 0
		maxCountCard := "J"
		for card, count := range cardsCount {
			if card == "J" {
				continue
			}

			if count > maxCount {
				maxCount = count
				maxCountCard = card
			}
		}

		cardsCount[maxCountCard] += cardsCount["J"]
		cardsCount["J"] = 0

	}

	previousCount := 0
	for _, count := range cardsCount {
		// five or four of a kind
		if count >= 4 {
			return int(math.Pow10(count))
		}

		if count >= 2 && previousCount == 0 {
			previousCount = count
			continue
		}

		// full house or two pairs
		if count >= 2 && previousCount > 0 {
			return int(math.Pow10(count)) + int(math.Pow10(previousCount))
		}
	}

	// three of a kind, one pair or high card
	return int(math.Pow10(previousCount))
}

func (h1 Hand) compare(h2 Hand, withJokers bool) int {
	// h1 >  h2 ->  1
	// h1 == h2 ->  0
	// h1 <  h2 -> -1
	if combDiff := h1.combinationStrength - h2.combinationStrength; combDiff != 0 {
		// why? because
		return int(math.Copysign(1, float64(combDiff)))
	}

	var cardValueMap map[string]int
	if withJokers {
		cardValueMap = map[string]int{
			"A": 14,
			"K": 13,
			"Q": 12,
			"T": 10,
			"9": 9,
			"8": 8,
			"7": 7,
			"6": 6,
			"5": 5,
			"4": 4,
			"3": 3,
			"2": 2,
			"J": 1,
		}
	} else {
		cardValueMap = map[string]int{
			"A": 14,
			"K": 13,
			"Q": 12,
			"J": 11,
			"T": 10,
			"9": 9,
			"8": 8,
			"7": 7,
			"6": 6,
			"5": 5,
			"4": 4,
			"3": 3,
			"2": 2,
		}
	}

	for i := 0; i < len(h1.cards); i++ {
		card1 := h1.cards[i]
		card2 := h2.cards[i]

		cardDiff := cardValueMap[card1] - cardValueMap[card2]
		if cardDiff != 0 {
			return int(math.Copysign(1, float64(cardDiff)))
		}
	}

	return 0
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLineSplit := strings.Split(scanner.Text(), " ")
		newHand := makeHand(
			strings.TrimSpace(currentLineSplit[0]),
			strings.TrimSpace(currentLineSplit[1]),
			false,
		)

		placed := false
		for i := 0; i < len(hands); i++ {
			if newHand.compare(hands[i], false) <= 0 {
				hands = slices.Insert(hands, i, newHand)
				placed = true
				break
			}
		}

		if !placed {
			hands = append(hands, newHand)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}

	return result
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLineSplit := strings.Split(scanner.Text(), " ")
		newHand := makeHand(
			strings.TrimSpace(currentLineSplit[0]),
			strings.TrimSpace(currentLineSplit[1]),
			true,
		)

		placed := false
		for i := 0; i < len(hands); i++ {
			if newHand.compare(hands[i], true) <= 0 {
				hands = slices.Insert(hands, i, newHand)
				placed = true
				break
			}
		}

		if !placed {
			hands = append(hands, newHand)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}

	return result
}
