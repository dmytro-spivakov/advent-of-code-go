package day02

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id       int
	maxRed   int
	maxGreen int
	maxBlue  int
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Failed to open input file for reading")
	}

	var games []game

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err = scanner.Err(); err != nil {
			log.Fatalf("Error during input file read: %v\n", err)
		}

		currentLine := scanner.Text()
		games = append(games, makeGame(currentLine))
	}

	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	result := 0
	for _, game := range games {
		if game.maxRed <= maxRed && game.maxGreen <= maxGreen && game.maxBlue <= maxBlue {
			result += game.id
		}
	}

	return result
}

func makeGame(gameLine string) game {
	gameIdRegexp := regexp.MustCompile(`Game (?<gameId>[0-9]+):`)
	gameIdString := gameIdRegexp.FindStringSubmatch(gameLine)[gameIdRegexp.SubexpIndex("gameId")]
	gameId := parseNumber(gameIdString)

	game := game{id: gameId}

	gameSetStrings := strings.Split(strings.Split(gameLine, ":")[1], ";")
	gameSetRegex := regexp.MustCompile(`(?<count>[0-9]+) (?<color>[a-z]+)`)

	for _, gameSetString := range gameSetStrings {
		for _, m := range gameSetRegex.FindAllStringSubmatch(gameSetString, -1) {
			color := m[gameSetRegex.SubexpIndex("color")]
			count := parseNumber(m[gameSetRegex.SubexpIndex("count")])

			switch color {
			case "red":
				if game.maxRed < count {
					game.maxRed = count
				}
			case "green":
				if game.maxGreen < count {
					game.maxGreen = count
				}
			case "blue":
				if game.maxBlue < count {
					game.maxBlue = count
				}
			default:
				log.Fatalf("Unknown color %v in game %v\n", color, game.id)
			}
		}
	}

	return game
}

func parseNumber(num string) int {
	number, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number: %v\n", num)
	}

	return int(number)
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Failed to open input file for reading")
	}

	var games []game

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err = scanner.Err(); err != nil {
			log.Fatalf("Error during input file read: %v\n", err)
		}

		currentLine := scanner.Text()
		games = append(games, makeGame(currentLine))
	}

	result := 0
	for _, game := range games {
		result += game.maxRed * game.maxGreen * game.maxBlue
	}

	return result
}
