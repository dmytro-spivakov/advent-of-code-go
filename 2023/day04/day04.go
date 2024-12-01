package day04

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers map[int]bool
	yourNumbers    map[int]bool
	copies         int
}

func (c Card) points() int {
	count := c.matchingNumbersCount()

	if count == 0 {
		return 0
	}
	return int(math.Pow(2, float64(count-1)))
}

func (c Card) matchingNumbersCount() int {
	count := 0
	for number := range c.yourNumbers {
		if c.winningNumbers[number] {
			count += 1
		}
	}

	return count
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Failed to open the input file")
	}

	var cards []Card
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cards = append(cards, parseCard(scanner.Text()))
	}
	if scanner.Err() != nil {
		log.Fatalln("Error during input file read")
	}

	result := 0
	for _, card := range cards {
		result += card.points()
	}
	return result
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Failed to open the input file")
	}

	cardsMap := make(map[int]Card)
	var orderedKeys []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := parseCard(scanner.Text())
		orderedKeys = append(orderedKeys, card.id)
		cardsMap[card.id] = card
	}
	if scanner.Err() != nil {
		log.Fatalln("Error during input file read")
	}

	result := 0
	for _, cardId := range orderedKeys {
		currentCard := cardsMap[cardId]
		result += currentCard.copies

		cardsToCopy := currentCard.matchingNumbersCount()
		currentCopyId := cardId + 1
		for cardsToCopy > 0 {
			if nextCard, ok := cardsMap[currentCopyId]; ok {
				nextCard.copies += currentCard.copies
				cardsMap[currentCopyId] = nextCard
			}

			currentCopyId += 1
			cardsToCopy -= 1
		}
	}
	return result
}

func parseCard(row string) Card {
	cardIdRegex := regexp.MustCompile(`Card\s+([0-9]+):`)
	stringId := cardIdRegex.FindStringSubmatch(row)[1]
	id := parseInt(stringId, "Card Id")

	numbersSubstring := strings.Split(strings.Split(row, ":")[1], "|")
	numberRegex := regexp.MustCompile(`([0-9]+)`)

	winningNumbers := make(map[int]bool)
	for _, numString := range numberRegex.FindAllString(numbersSubstring[0], -1) {
		winningNumbers[parseInt(numString, "Winning number")] = true
	}

	yourNumbers := make(map[int]bool)
	for _, numString := range numberRegex.FindAllString(numbersSubstring[1], -1) {
		yourNumbers[parseInt(numString, "Your number")] = true
	}

	return Card{id: id, winningNumbers: winningNumbers, yourNumbers: yourNumbers, copies: 1}
}

func parseInt(num string, errLocation string) int {
	intNum, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse %v with value %v\n", errLocation, num)
	}

	return int(intNum)
}
