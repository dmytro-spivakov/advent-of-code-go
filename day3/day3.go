/**
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

--- Part Two ---

The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

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
In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
**/

package day3

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type numberWithCoordinates struct {
	y      int
	startX int
	endX   int
	number int
}

type asteriscCoordinates struct {
	y int
	x int
}

var numberRegex = regexp.MustCompile(`[0-9]{1}`)

func Solution() int {
	inputFile, err := os.Open("./day3/input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inputFile)
	var charsMatrix [][]string
	for scanner.Scan() {
		currentLine := scanner.Text()
		charsMatrix = append(charsMatrix, strings.Split(currentLine, ""))
	}

	sum := 0

	numberCoordsSlice := []numberWithCoordinates{}
	asteriscCoordinatesSlice := []asteriscCoordinates{}
	for y, row := range charsMatrix {
		for x, rowItem := range row {
			if rowItem == "*" {
				asteriscCoordinatesSlice = append(asteriscCoordinatesSlice, asteriscCoordinates{x: x, y: y})
				continue
			} else if !numberRegex.Match([]byte(rowItem)) {
				continue
			}

			newNumber := getNumber(charsMatrix, x, y)
			duplicate := false
			for _, visitedNumber := range numberCoordsSlice {
				if newNumber.y == visitedNumber.y && newNumber.startX == visitedNumber.startX && newNumber.number == visitedNumber.number {
					duplicate = true
				}
			}

			if !duplicate {
				numberCoordsSlice = append(numberCoordsSlice, newNumber)
			}
		}
	}

	for _, asteriscCoordsPair := range asteriscCoordinatesSlice {
		adjacentNumberCoordsSlice := []numberWithCoordinates{}

		for _, numberCoords := range numberCoordsSlice {
			// asterisc y and x are astX and astY
			// the number is adjacent to the asterisc if
			// - the number is within charsMatrix[astY - 1][astX - 1...astX + 1]
			// - the number is within charsMatrix[astY][astX - 1...astX + 1]
			// - the number is within charsMatrix[astY + 1][astX - 1...astX + 1]

			for _, astY := range [3]int{asteriscCoordsPair.y - 1, asteriscCoordsPair.y, asteriscCoordsPair.y + 1} {
				if astY < 0 || astY > len(charsMatrix)-1 {
					continue
				}

				if numberCoords.y == astY && numberCoords.startX <= asteriscCoordsPair.x+1 && numberCoords.endX >= asteriscCoordsPair.x-1 {
					adjacentNumberCoordsSlice = append(adjacentNumberCoordsSlice, numberCoords)
				}

			}
		}

		if len(adjacentNumberCoordsSlice) != 2 {
			continue
		}

		adjacentNumbers := [2]int{}

		for i, adjacentNumberCoords := range adjacentNumberCoordsSlice {
			numberString := strings.Join(charsMatrix[adjacentNumberCoords.y][adjacentNumberCoords.startX:adjacentNumberCoords.endX+1], "")
			number, err := strconv.ParseInt(numberString, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			adjacentNumbers[i] = int(number)
		}

		sum += adjacentNumbers[0] * adjacentNumbers[1]

	}

	err = inputFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	return sum
}

func getNumber(matrix [][]string, x int, y int) numberWithCoordinates {
	matrixRow := matrix[y]

	startX := x
	for startX >= 0 && numberRegex.Match([]byte(matrixRow[startX])) {
		startX -= 1
	}
	startX += 1

	endX := x
	for endX <= len(matrixRow)-1 && numberRegex.Match([]byte(matrixRow[endX])) {
		endX += 1
	}
	endX -= 1

	numberString := strings.Join(matrix[y][startX:endX+1], "")
	number, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return numberWithCoordinates{y: y, startX: startX, endX: endX, number: int(number)}
}
