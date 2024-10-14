package day03

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type number struct {
	value int
	x1    int
	x2    int
	y     int
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error opening the input file")
	}

	var numbers []number
	specialSymbolCoords := make(map[int]map[int]bool)

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			log.Fatalf("Error reading the input file: %v\n", err)
		}

		rowNumbers, rowSpecialSymbols := extractRow(strings.Split(scanner.Text(), ""), y)
		numbers = append(numbers, rowNumbers...)
		specialSymbolCoords[y] = rowSpecialSymbols
		y += 1
	}

	/*
		Walk through all the numbers, add to the result sum if a special symbol with adjacent coords exists.
		adjacent coords: x = [x1 - 1...x2 + 1], y = [y - 1...y + 1], walk x * y
	*/
	result := 0
	for _, num := range numbers {
		alreadyFound := false

		for y := num.y - 1; y <= num.y+1; y++ {
			for x := num.x1 - 1; x <= num.x2+1; x++ {
				if specialSymbolCoords[y][x] && !alreadyFound {
					result += num.value
					// goto's are spooky, use a flag to prevent dup entries of the same number
					alreadyFound = true
				}
			}
		}
	}

	return result
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error opening the input file")
	}

	var numbers []number
	specialSymbolCoords := make(map[int]map[int]bool)

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			log.Fatalf("Error reading the input file: %v\n", err)
		}

		rowNumbers, rowSpecialSymbols := extractRow(strings.Split(scanner.Text(), ""), y)
		numbers = append(numbers, rowNumbers...)
		specialSymbolCoords[y] = rowSpecialSymbols
		y += 1
	}

	result := 0
	for y, specialSymbolsRow := range specialSymbolCoords {
		for x := range specialSymbolsRow {
			adjacentNumbers := []int{}

			for _, number := range numbers {
				if (number.y >= y-1 && number.y <= y+1) && (number.x1 <= x+1 && number.x2 >= x-1) {
					adjacentNumbers = append(adjacentNumbers, number.value)
				}
			}

			if len(adjacentNumbers) == 2 {
				result += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}

	return result
}

func extractRow(row []string, currentY int) (numbers []number, specialSymbolCoords map[int]bool) {
	digitRegexp := regexp.MustCompile(`^[0-9]{1}$`)
	specialSymbolCoords = map[int]bool{}

	currentNumber := ""
	for x, char := range row {
		if digitRegexp.MatchString(char) {
			currentNumber += char
		}

		if !digitRegexp.MatchString(char) && char != "." {
			specialSymbolCoords[x] = true
		}

		if len(currentNumber) > 0 {
			/*
				first case: next char is anything but a digit - we've encountered the end of the current number.
				Pop `currentNumber` (poor man's stack of 1) and turn the concatenated digits string into number.

				second case: handle the edge-case when the number is at the right edge of the line.
			*/
			if !digitRegexp.MatchString(char) {
				intNumber, err := strconv.ParseInt(currentNumber, 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse number %v with %v", currentNumber, err)

				}

				numbers = append(numbers, number{value: int(intNumber), y: currentY, x2: x - 1, x1: x - len(currentNumber)})
				currentNumber = ""

			} else if x >= len(row)-1 {
				intNumber, err := strconv.ParseInt(currentNumber, 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse number %v with %v", currentNumber, err)

				}

				numbers = append(numbers, number{value: int(intNumber), y: currentY, x2: x, x1: x + 1 - len(currentNumber)})
				currentNumber = ""
			}
		}

	}
	return numbers, specialSymbolCoords
}
