package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Solution1(filepath string) int {
	matrix := parseInputMatrix(filepath, true)
	// printMatrix(matrix)
	galaxies := findGalaxies(&matrix)
	gPairs := makePairs(galaxies)

	result := 0
	for _, pair := range gPairs {
		res := naiveDistCalc(pair[0], pair[1], pair[2], pair[3])
		result += res
	}
	return result
}

func Solution2(filepath string) int {
	matrix := parseInputMatrix(filepath, false)
	// printMatrix(matrix)
	galaxies := findGalaxies(&matrix)
	gPairs := makePairs(galaxies)

	result := 0
	dupVertical, dupHorizontal := findBlankLines(matrix)
	for _, pair := range gPairs {
		res := naiveDistCalc2(pair[0], pair[1], pair[2], pair[3], dupVertical, dupHorizontal)
		result += res
	}
	return result
}

func naiveDistCalc2(y1, x1, y2, x2 int, dupVertical, dupHorizontal []int) int {
	dist := 0

	for x1 != x2 {
		diff := 1
		if _, found := slices.BinarySearch(dupVertical, x1); found {
			diff = 1000000
		}

		if x1 > x2 {
			x1--
		} else {
			x1++
		}
		dist += diff
	}

	for y1 != y2 {
		diff := 1
		if _, found := slices.BinarySearch(dupHorizontal, y1); found {
			diff = 1000000
		}

		if y1 > y2 {
			y1--
		} else {
			y1++
		}
		dist += diff
	}

	return dist
}

func naiveDistCalc(y1, x1, y2, x2 int) int {
	dist := 0

	for x1 != x2 {
		if x1 > x2 {
			x1--
		} else {
			x1++
		}
		dist++
	}

	for y1 != y2 {
		if y1 > y2 {
			y1--
		} else {
			y1++
		}
		dist++
	}

	return dist
}

func makePairs(galaxies [][2]int) [][4]int {
	var coordPairs [][4]int

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			newPair := [4]int{
				galaxies[i][0],
				galaxies[i][1],
				galaxies[j][0],
				galaxies[j][1],
			}
			coordPairs = append(coordPairs, newPair)
		}
	}

	return coordPairs
}

func findGalaxies(matrix *[][]string) [][2]int {
	m := *matrix
	var gCoords [][2]int

	for y, row := range m {
		for x, el := range row {
			if el == "#" {
				gCoords = append(gCoords, [2]int{y, x})
			}
		}
	}

	return gCoords
}

func parseInputMatrix(filepath string, normalize bool) [][]string {
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
		log.Fatalf("Error during reading the input file: %v\n", err.Error())
	}

	if normalize {
		return normalizeMatrix(matrix)
	} else {
		return matrix
	}
}

func normalizeMatrix(matrix [][]string) [][]string {
	var normalized [][]string

	dupVertical, dupHorizontal := findBlankLines(matrix)

	for y := len(matrix) - 1; y >= 0; y-- {
		newRow := make([]string, len(matrix[0]), len(matrix[0])+len(dupVertical))
		copy(newRow, matrix[y])
		for x := len(dupVertical) - 1; x >= 0; x-- {
			newRow = slices.Insert(newRow, dupVertical[x], ".")
		}

		normalized = slices.Insert(normalized, 0, newRow)
		if _, found := slices.BinarySearch(dupHorizontal, y); found {
			normalized = slices.Insert(normalized, 0, newRow)
		}
	}

	return normalized
}

func findBlankLines(matrix [][]string) ([]int, []int) {
	var dupVertical, dupHorizontal []int
	for x := 0; x < len(matrix[0]); x++ {
		dup := true
		for y := 0; y < len(matrix); y++ {
			if matrix[y][x] == "#" {
				dup = false
			}
		}

		if dup {
			dupVertical = append(dupVertical, x)
		}
	}
	for y, row := range matrix {
		if !slices.Contains(row, "#") {
			dupHorizontal = append(dupHorizontal, y)
		}
	}
	slices.Sort(dupVertical)
	slices.Sort(dupHorizontal)

	return dupVertical, dupHorizontal
}

func printMatrix(matrix [][]string) {
	fmt.Println("----------MATRIX START----------")
	for _, mRow := range matrix {
		fmt.Println(mRow)
	}
	fmt.Println("-----------MATRIX END----------")
}
