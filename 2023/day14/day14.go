package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var m [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, strings.Split(scanner.Text(), ""))
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	m = tiltNorth(m)
	return calcWeight(m)
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var m [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, strings.Split(scanner.Text(), ""))
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return calcFinalWeight(m)
}

func calcWeight(m [][]string) int {
	result := 0
	for y := len(m) - 1; y >= 0; y-- {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == "O" {
				result += len(m) - y
			}
		}

	}

	return result
}

// idea:
// implement the base 4 direction roll algorithm,
// then add logic to detect number of iterations required for the pattern to repeat (?ignore the initial state where boulders are placed randomly)
// once a repat is detected skip to (1000000000 - current number of cycles) % repeat length
func calcFinalWeight(m [][]string) int {
	remaining := 1000000000

	// { serializedMatrix: loop idx, ... }
	loopMap := make(map[string]int)
	var loopStart, loopEnd int

	for i := 1; i < 1000; i++ {
		for j := 0; j <= 3; j++ {
			m = tiltNorth(m)
			m = rotateClockwise(m)
		}

		remaining--
		currentCycle := fmt.Sprintf("%v", m)
		if idx, ok := loopMap[currentCycle]; ok {
			loopStart, loopEnd = idx, i
			break
		} else {
			loopMap[currentCycle] = i
		}
	}

	loopLength := loopEnd - loopStart
	remaining = remaining % loopLength

	result := 0
	for remaining > 0 {
		for j := 0; j <= 3; j++ {
			m = tiltNorth(m)
			m = rotateClockwise(m)
		}
		result = calcWeight(m)
		remaining--
	}

	return result
}

func rotateClockwise(m [][]string) [][]string {
	transposed := make([][]string, len(m[0]))

	for x := 0; x < len(m[0]); x++ {
		for y := len(m) - 1; y >= 0; y-- {
			transposed[x] = append(transposed[x], m[y][x])
		}
	}

	return transposed
}

func tiltNorth(m [][]string) [][]string {
	availablePos := make(map[int]int)

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			char := m[y][x]
			if char == "#" {
				availablePos[x] = y + 1
				continue
			}
			if char == "O" {
				newY := availablePos[x]
				if y != newY && newY < len(m) {
					m[newY][x] = "O"
					m[y][x] = "."
				}
				availablePos[x] = availablePos[x] + 1
			}
		}
	}

	return m
}

func printMatrix(m [][]string) {
	fmt.Println("-----MATRIX START-----")
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println("------MATRIX END------")
}
