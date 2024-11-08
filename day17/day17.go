package day17

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PQ struct {
	queue [][6]int
}

func (pq *PQ) len() int {
	return len((*pq).queue)
}
func (pq *PQ) pop() (int, int, int, int, int, int) {
	pq.sort()
	r := (*pq).queue[0]
	(*pq).queue = (*pq).queue[1:]

	return r[0], r[1], r[2], r[3], r[4], r[5]
}
func (pq *PQ) push(path [6]int) {
	(*pq).queue = append((*pq).queue, path)
}
func (pq *PQ) sort() {
	slices.SortFunc((*pq).queue, func(a, b [6]int) int {
		if a[0] > b[0] {
			return 1
		} else if a[0] < b[0] {
			return -1
		} else {
			return 0
		}
	})
}

func Solution1(filepath string) int {
	m := readInput(filepath)

	pq := PQ{}
	seen := make(map[string]bool)
	pq.push([6]int{0, 0, 0, 0, 0, 0}) // heat loss, y, x, dY, dX, n of steps in straight line

	for pq.len() > 0 {
		heatLoss, y, x, dY, dX, n := pq.pop()
		if y == len(m)-1 && x == len(m[y])-1 {
			return heatLoss
		}

		seenKey := fmt.Sprintf("%d;%d;%d;%d;%d", y, x, dY, dX, n)
		if seen[seenKey] {
			continue
		}
		seen[seenKey] = true

		// keep going in the same direction
		if n < 3 && [2]int{dY, dX} != [2]int{0, 0} {
			newY, newX := y+dY, x+dX
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[0]) {
				// do nothing, I don't want to invert this condition
			} else {
				pq.push([6]int{heatLoss + m[newY][newX], newY, newX, dY, dX, n + 1})
			}
		}

		// explore all the other directions except for the current and its reverse
		for _, diffs := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newDY, newDX := diffs[0], diffs[1]
			newY, newX := y+newDY, x+newDX
			if newDir := [2]int{newDY, newDX}; newDir == [2]int{dY, dX} || newDir == [2]int{-dY, -dX} {
				continue
			}
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[newY]) {
				continue
			}
			pq.push([6]int{heatLoss + m[newY][newX], newY, newX, newDY, newDX, 1})
		}

	}
	return -1
}

func Solution2(filepath string) int {
	m := readInput(filepath)

	pq := PQ{}
	seen := make(map[string]bool)
	pq.push([6]int{0, 0, 0, 0, 0, 0}) // heat loss, y, x, dY, dX, n of steps in straight line

	for pq.len() > 0 {
		heatLoss, y, x, dY, dX, n := pq.pop()
		// fmt.Printf("DEBUG: hl=%d, y=%d, x=%d, dY=%d, dX=%d, n=%d\n", heatLoss, y, x, dY, dX, n)
		if y == len(m)-1 && x == len(m[y])-1 && n >= 4 {
			return heatLoss
		}

		seenKey := fmt.Sprintf("%d;%d;%d;%d;%d", y, x, dY, dX, n)
		if seen[seenKey] {
			continue
		}
		seen[seenKey] = true

		// keep going in the same direction
		if n < 10 && [2]int{dY, dX} != [2]int{0, 0} {
			newY, newX := y+dY, x+dX
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[0]) {
				// do nothing, I don't want to invert this condition
			} else {
				pq.push([6]int{heatLoss + m[newY][newX], newY, newX, dY, dX, n + 1})
			}
		}

		if n != 0 && n < 4 {
			continue
		}
		// explore all the other directions except for the current and its reverse
		for _, diffs := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newDY, newDX := diffs[0], diffs[1]
			newY, newX := y+newDY, x+newDX
			if newDir := [2]int{newDY, newDX}; newDir == [2]int{dY, dX} || newDir == [2]int{-dY, -dX} {
				continue
			}
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[newY]) {
				continue
			}
			pq.push([6]int{heatLoss + m[newY][newX], newY, newX, newDY, newDX, 1})
		}

	}
	return -1
}

func readInput(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var matrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentRow := strings.Split(scanner.Text(), "")
		currentRowNums := make([]int, len(currentRow))
		for i, el := range currentRow {
			currentRowNums[i] = parseInt(el)
		}
		matrix = append(matrix, currentRowNums)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return matrix
}

func parseInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with: %s\n", s, err.Error())
	}

	return int(num)
}

func printMatrix(m [][]int) {
	fmt.Println("-----MATRIX START-----")
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println("------MATRIX END------")
}
