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
	matrix := parseInputMatrix(filepath)
	printMatrix(matrix)
	galaxies := findGalaxies(&matrix)
	fmt.Printf("Galaxies %d: %v\n", len(galaxies), galaxies)
	gPairs := makePairs(galaxies)
	fmt.Printf("All pairs: %v\n", gPairs)

	result := 0
	for _, pair := range gPairs {
		res := naiveDistCalc(pair[0], pair[1], pair[2], pair[3])
		// fmt.Printf("ADDING {%d, %d}, {%d, %d} = %d\n", pair[0], pair[1], pair[2], pair[3], res)
		result += res
	}
	return result
}

func Solution2(filepath string) int {
	return 0
}

func naiveDistCalc(x1, y1, x2, y2 int) int {
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

func parseInputMatrix(filepath string) [][]string {
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

	return normalizeMatrix(matrix)
}

func normalizeMatrix(matrix [][]string) [][]string {
	var normalized [][]string

	var dVertical, dHorizonatl []int
	for x := 0; x < len(matrix[0]); x++ {
		dup := true
		for y := 0; y < len(matrix); y++ {
			if matrix[y][x] == "#" {
				dup = false
			}
		}

		if dup {
			dVertical = append(dVertical, x)
		}
	}
	for y, row := range matrix {
		if !slices.Contains(row, "#") {
			dHorizonatl = append(dHorizonatl, y)
		}
	}
	slices.Sort(dVertical)
	slices.Sort(dHorizonatl)

	for y := len(matrix) - 1; y >= 0; y-- {
		newRow := make([]string, len(matrix[0]), len(matrix[0])+len(dVertical))
		copy(newRow, matrix[y])
		for x := len(dVertical) - 1; x >= 0; x-- {
			newRow = slices.Insert(newRow, dVertical[x], ".")
		}

		normalized = slices.Insert(normalized, 0, newRow)
		if _, found := slices.BinarySearch(dHorizonatl, y); found {
			normalized = slices.Insert(normalized, 0, newRow)
		}
	}

	return normalized
}

func printMatrix(matrix [][]string) {
	fmt.Println("----------MATRIX START----------")
	for _, mRow := range matrix {
		fmt.Println(mRow)
	}
	fmt.Println("-----------MATRIX END----------")
}
