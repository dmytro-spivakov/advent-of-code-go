package day14

import (
	"bufio"
	"log"
	"os"
)

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var results []int
	// { x: y }
	maxAvailablePositon := make(map[int]int)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, char := range scanner.Text() {
			switch char {
			case 'O':
				results = append(results, maxAvailablePositon[x])
				maxAvailablePositon[x] = maxAvailablePositon[x] + 1
			case '#':
				maxAvailablePositon[x] = y + 1
			}
		}

		y++
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	result := 0
	maxY := y
	for _, r := range results {
		result += maxY - r
	}

	return result
}

func Solution2(filepath string) int {
	return -1
}
