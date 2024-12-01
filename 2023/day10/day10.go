package day10

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Solution1(filepath string) int {
	inputMatrix, start := readMatrix(filepath)

	visited := make(map[int]map[int]bool)
	var results [][][2]int
	normalizeMatrix(&inputMatrix, start)
	findLoops(&inputMatrix, start, start, &visited, make([][2]int, 0), &results)

	maxLen := 0
	for _, res := range results {
		if curLen := len(res); curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen / 2
}

func readMatrix(filepath string) (inputMatrix [][]string, startPos [2]int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		startIndex := slices.Index(row, "S")
		if startIndex >= 0 {
			startPos = [2]int{len(inputMatrix), startIndex}
		}
		inputMatrix = append(inputMatrix, row)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return inputMatrix, startPos
}

func findLoops(inputMatrix *[][]string, current [2]int, end [2]int, visited *map[int]map[int]bool, currentPath [][2]int, results *[][][2]int) {
	curY, curX := current[0], current[1]
	currentPath = append(currentPath, current)

	// base case 1: found S
	// check visited not to stop search on the first call (starting from S)
	if curY == end[0] && curX == end[1] && (*visited)[curY][curX] {
		*results = append(*results, currentPath)
		return
	}
	if (*visited)[curY] == nil {
		(*visited)[curY] = make(map[int]bool)
	}
	(*visited)[curY][curX] = true

	// recursive flow
	allAdjacent := findAdjacent(inputMatrix, current)
	var viableAdjacent [][2]int
	for _, adjacent := range allAdjacent {
		if !(*visited)[adjacent[0]][adjacent[1]] || (adjacent[0] == end[0] && adjacent[1] == end[1]) {
			viableAdjacent = append(viableAdjacent, adjacent)
		}
	}

	for _, adjacent := range viableAdjacent {
		findLoops(inputMatrix, adjacent, end, visited, currentPath, results)
	}
}

func normalizeMatrix(inputMatrix *[][]string, current [2]int) {
	adjacent := make([][2]int, 0)

	if (*inputMatrix)[current[0]][current[1]] == "S" {
		for y := current[0] - 1; y < len(*inputMatrix); y++ {
			if y < 0 {
				continue
			}
			for x := current[1] - 1; x < len((*inputMatrix)[0]); x++ {
				if x < 0 || (y == current[0] && x == current[1]) {
					continue
				}
				neighbourAdj := findAdjacent(inputMatrix, [2]int{y, x})
				if slices.Contains(neighbourAdj, current) {
					adjacent = append(adjacent, [2]int{y, x})
				}
			}
		}
	}

	startingChar := ""
	for char, coordDiff := range legendMap() {
		reverseNavCoords := [][2]int{
			{adjacent[0][0] - coordDiff[0][0], adjacent[0][1] - coordDiff[0][1]},
			{adjacent[1][0] - coordDiff[1][0], adjacent[1][1] - coordDiff[1][1]},
		}

		if reverseNavCoords[0] == reverseNavCoords[1] && reverseNavCoords[0] == current {
			startingChar = char
			break
		}
	}
	(*inputMatrix)[current[0]][current[1]] = startingChar
}

func findAdjacent(inputMatrix *[][]string, current [2]int) [][2]int {
	adjacent := make([][2]int, 0)

	for char, targetCoords := range legendMap() {
		if (*inputMatrix)[current[0]][current[1]] == char {
			adjacent = append(
				adjacent,
				[2]int{current[0] + targetCoords[0][0], current[1] + targetCoords[0][1]},
				[2]int{current[0] + targetCoords[1][0], current[1] + targetCoords[1][1]},
			)
		}
	}

	return adjacent
}

func legendMap() map[string][][2]int {
	return map[string][][2]int{
		"-": {{0, -1}, {0, 1}},
		"|": {{-1, 0}, {1, 0}},
		"L": {{-1, 0}, {0, 1}},
		"F": {{0, 1}, {1, 0}},
		"J": {{-1, 0}, {0, -1}},
		"7": {{0, -1}, {1, 0}},
	}
}

func printMatrix(matrix [][]string) {
	// fmt.Println("---------------------------")
	// for _, row := range matrix {
	// 	fmt.Println(strings.Join(row, " "))
	// }
	// fmt.Println("---------------------------")
}

func Solution2(filepath string) int {
	inputMatrix, start := readMatrix(filepath)

	visited := make(map[int]map[int]bool)
	var results [][][2]int
	normalizeMatrix(&inputMatrix, start)
	findLoops(&inputMatrix, start, start, &visited, make([][2]int, 0), &results)

	var mainLoop [][2]int
	for _, res := range results {
		if len(res) > len(mainLoop) {
			mainLoop = res
		}
	}
	// fmt.Printf("MAIN LOOP: %v\n", mainLoop)

	scannedMatrix := scanMatrix(inputMatrix, mainLoop)
	// fmt.Println("SCANNED MATRIX:")
	printMatrix(scannedMatrix)

	result := 0
	for _, scannedRow := range scannedMatrix {
		for _, scannedChar := range scannedRow {
			if scannedChar == "x" {
				result++
			}
		}
	}

	return result
}

func scanMatrix(inputMatrinx [][]string, mainLoop [][2]int) [][]string {
	var scannedMatrix [][]string
	for _, inputRow := range inputMatrinx {
		copiedRow := make([]string, len(inputRow))
		copy(copiedRow, inputRow)
		scannedMatrix = append(scannedMatrix, copiedRow)
	}

	for y := 0; y < len(scannedMatrix); y++ {
		withinLoop := false
		previousBend := ""
		for x := 0; x < len(scannedMatrix[0]); x++ {
			current := inputMatrinx[y][x]
			loopPipe := slices.Contains(mainLoop, [2]int{y, x})

			// anything that isn't a part of the main loop,
			// it doesn't matter if it's ground or pipe
			if !loopPipe {
				if withinLoop {
					scannedMatrix[y][x] = "x"
				} else {
					scannedMatrix[y][x] = "o"
				}
				continue
			}

			// handle parts of the main loop
			// easy case:
			// whenever we ecounter "|" we cross the loop
			// if the point is inside the loop we'd always cross the loop odd number of times
			// if its outside - even number of times
			// therefore bool toggle, counting and doing % == 0 would work too.
			if current == "|" {
				withinLoop = !withinLoop
				previousBend = ""
			}

			// the fact that it's valid properly connected loop is assured by the part 1
			// we know that we're currently in the loop and the loop was discovered by recursively traversing it
			// i.e. there's no reason to care for fringe cases such F + J
			//
			// treat bends such as L-*J, F-*7 as 1 loop crossing
			if current == "L" || current == "F" {
				withinLoop = !withinLoop
				previousBend = current
			}
			if current == "J" || current == "7" {
				if (current == "J" && previousBend == "L") || (current == "7" && previousBend == "F") {
					withinLoop = !withinLoop
					previousBend = ""
				}
			}

		}
	}

	return scannedMatrix
}

func toggleBool(b bool) bool {
	if b {
		return false
	} else {
		return true
	}
}
