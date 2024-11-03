package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solution1(filepath string) int {
	m := readInput(filepath)
	fmt.Println(m)

	// "0,0": true
	energizedCells := make(map[string]bool)

	rayQueue := [][4]int{{0, 0, 0, 1}}
	// "0,0;0,1": true where position is encoded before ; and direction is after
	processedRays := make(map[string]bool)

	for len(rayQueue) > 0 {
		currentRay := rayQueue[0]
		rayQueue = rayQueue[1:]

		y, x := currentRay[0], currentRay[1]
		diffY, diffX := currentRay[2], currentRay[3]
		rayKey := fmt.Sprintf("%d,%d;%d,%d", y, x, diffY, diffX)
		if _, ok := processedRays[rayKey]; ok {
			// junction with the same position and ray direction has already been visited
			continue
		} else {
			processedRays[rayKey] = true
		}

		for {
			if x < 0 || x >= len(m[0]) || y < 0 || y >= len(m) {
				break // ray left the plane
			}
			energizedCells[fmt.Sprintf("%d,%d", y, x)] = true
			fmt.Printf("DEBUG: pos = %d,%d; diff = %d,%d\n", y, x, diffY, diffX)

			rayEnd := false
			switch m[y][x] {
			case ".":
			case "\\":
				// reflect == swap diffX and diffY
				// ???
				// rayQueue = append(rayQueue, [4]int{y, x, diffX, diffY})
				// rayEnd = true
				diffY, diffX = diffX, diffY
			case "/":
				// rayQueue = append(rayQueue, [4]int{y, x, -diffX, -diffY})
				// rayEnd = true
				diffY, diffX = -diffX, -diffY
			case "-":
				if diffY != 0 && diffX == 0 {
					rayQueue = append(rayQueue, [4]int{y, x - 1, 0, -1})
					rayQueue = append(rayQueue, [4]int{y, x + 1, 0, 1})
					rayEnd = true
				}
			case "|":
				if diffX != 0 && diffY == 0 {
					rayQueue = append(rayQueue, [4]int{y + 1, x, 1, 0})
					rayQueue = append(rayQueue, [4]int{y - 1, x, -1, 0})
					rayEnd = true
				}
			}

			if rayEnd {
				break
			}

			// it's either misaligned split or a dot
			y += diffY
			x += diffX
		}
	}

	printMatrix(m)
	res := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if _, ok := energizedCells[fmt.Sprintf("%d,%d", y, x)]; ok {
				m[y][x] = "#"
				res++
			} else {
				m[y][x] = "."
			}
		}
	}
	printMatrix(m)
	fmt.Printf("DEBUG: res=%d\n", res)
	return len(energizedCells)
}

func Solution2(filepath string) int {
	return -1
}

func readInput(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var matrix [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, strings.Split(scanner.Text(), ""))
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return matrix
}

func printMatrix(m [][]string) {
	fmt.Println("-----MATRIX START-----")
	for _, row := range m {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println("------MATRIX END------")
}
