package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	result := 0
	scanner := bufio.NewScanner(file)
	var currentMatrix []string
	for scanner.Scan() {
		currentLine := scanner.Text()

		if len(currentLine) == 0 {
			result += calcReflection(currentMatrix)
			currentMatrix = make([]string, 0)
			continue
		}

		currentMatrix = append(currentMatrix, currentLine)
	}
	if len(currentMatrix) != 0 {
		result += calcReflection(currentMatrix)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return result
}

func Solution2(filepath string) int {
	return -1
}

func calcReflection(m []string) int {
	// vertical reflection line
	for i := 1; i < len(m); i++ {
		step := 0
		for {
			// imaginary reflection line is between upIdx=i-1 and downIdx=i, then we start with offset 0 and just keep increasing it by 1 on each iteration
			// we keep walking until:
			// a) success - we've reached one of the ends => return
			// b) failure - we've encountered m[upIdx] != m[downIdx] => break and move onto the next i
			upIdx := i - 1 - step
			downIdx := i + step

			if upIdx < 0 || downIdx >= len(m) {
				// i == number of rows above the current i
				separator := strings.Repeat("-", len(m[0]))
				m = slices.Insert(m, i, separator)
				printMatrix(m)
				return 100 * i
			}

			if m[upIdx] != m[downIdx] {
				fmt.Printf("i=%d, step=%d: %s != %s\n", i, step, m[upIdx], m[downIdx])
				break
			}
			step++
		}
	}

	// horizontal reflection line
	for i := 1; i < len(m[0]); i++ {
		step := 0
		for {
			leftIdx := i - 1 - step
			rightIdx := i + step

			if leftIdx < 0 || rightIdx >= len(m[0]) {
				separator := "|"
				for j := 0; j < len(m); j++ {
					m[j] = m[j][:i] + separator + m[j][i:]
				}
				printMatrix(m)
				// i == number of cols left of i
				return i
			}

			reflection := true
			for j := 0; j < len(m); j++ {
				currentRow := m[j]
				if currentRow[leftIdx] != currentRow[rightIdx] {
					reflection = false
					break
				}
			}
			if !reflection {
				break
			}
			step++
		}
	}

	return 0
}

func printMatrix(m []string) {
	fmt.Println("-----MATRIX START-----")
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println("------MATRIX END------")
}
