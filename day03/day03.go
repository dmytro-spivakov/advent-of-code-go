package day03

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 3: Gear Ratios ---

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

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
	specialSymbolsCoords := make(map[int]map[int]bool)
	y := 0
	scanner := bufio.NewScanner(file)
	digitRegexp := regexp.MustCompile(`^[0-9]{1}$`)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			log.Fatalf("Error reading the input file: %v\n", err)
		}

		specialSymbolsCoords[y] = make(map[int]bool)

		currentLine := strings.Split(scanner.Text(), "")
		currentNumber := ""
		for x, char := range currentLine {
			if digitRegexp.MatchString(char) {
				currentNumber += char
			}

			if !digitRegexp.MatchString(char) && char != "." {
				specialSymbolsCoords[y][x] = true
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

					numbers = append(numbers, number{value: int(intNumber), y: y, x2: x - 1, x1: x - len(currentNumber)})
					currentNumber = ""

				} else if x >= len(currentLine)-1 {
					intNumber, err := strconv.ParseInt(currentNumber, 10, 64)
					if err != nil {
						log.Fatalf("Failed to parse number %v with %v", currentNumber, err)

					}

					numbers = append(numbers, number{value: int(intNumber), y: y, x2: x, x1: x + 1 - len(currentNumber)})
					currentNumber = ""
				}
			}

		}

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
				if specialSymbolsCoords[y][x] && !alreadyFound {
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
	return 0
}
