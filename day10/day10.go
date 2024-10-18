package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Solution1(filepath string) int {
	inputMatrix, start := readMatrix(filepath)
	fmt.Printf("Starting at {%d, %d}\n", start[0], start[1])
	printMatrix(inputMatrix)

	visited := make(map[int]map[int]bool)
	var results [][][2]int
	findLoops(&inputMatrix, start, start, &visited, make([][2]int, 0), &results)

	fmt.Println("Results:")
	for _, res := range results {
		fmt.Println(res)
	}
	fmt.Println("-------------------")

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

func findAdjacent(inputMatrix *[][]string, current [2]int) [][2]int {
	adjacent := make([][2]int, 0)

	// handle unknown S
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
		return adjacent
	}

	// base case
	legendMap := map[string][][2]int{
		"-": {{0, -1}, {0, 1}},
		"|": {{-1, 0}, {1, 0}},
		"L": {{-1, 0}, {0, 1}},
		"F": {{0, 1}, {1, 0}},
		"J": {{-1, 0}, {0, -1}},
		"7": {{0, -1}, {1, 0}},
	}

	for char, targetCoords := range legendMap {
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

func printMatrix(matrix [][]string) {
	fmt.Println("---------------------------")
	for _, row := range matrix {
		fmt.Println(strings.Join(row, " "))
	}
	fmt.Println("---------------------------")
}

func Solution2(filepath string) int {
	return 0
}
