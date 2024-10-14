package day09

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}

	var readingsMatrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numberRegex := regexp.MustCompile(`[-]?\d+`)

		var currentRow []int
		for _, num := range numberRegex.FindAllString(scanner.Text(), -1) {
			currentRow = append(currentRow, parseInt(num))
		}
		readingsMatrix = append(readingsMatrix, currentRow)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	result := 0
	for _, matrixRow := range readingsMatrix {
		result += predictNextValue(matrixRow)
	}

	return result
}

func predictNextValue(matrixRow []int) int {
	// Each new row has x - 1 elements since x11 = x2 - x1, x12 = x3 - x2, ...
	// x x x x x
	// x x x x
	// x x x
	// x x
	// x
	predictionMaxtrix := makePredictionMatrix(matrixRow)
	// fmt.Println("PREDICT MATRIX:")
	// for i, newMatrixRow := range predictionMaxtrix {
	// 	fmt.Printf("%d: %v\n", i, newMatrixRow)
	// }
	// fmt.Println("----------------------------")

	for i := len(predictionMaxtrix) - 1; i >= 0; i-- {
		if i == len(predictionMaxtrix)-1 {
			predictionMaxtrix[i] = append(predictionMaxtrix[i], 0)
			continue
		}

		currenRowLastEl := predictionMaxtrix[i][len(predictionMaxtrix[i])-1]
		lowerRowLastEl := predictionMaxtrix[i+1][len(predictionMaxtrix[i+1])-1]
		predictionMaxtrix[i] = append(predictionMaxtrix[i], currenRowLastEl+lowerRowLastEl)
	}

	topRow := predictionMaxtrix[0]
	return topRow[len(topRow)-1]
}

func makePredictionMatrix(matrixRow []int) [][]int {
	allZeroes := true
	for _, el := range matrixRow {
		if el != 0 {
			allZeroes = false
		}
	}

	result := [][]int{matrixRow}
	if allZeroes {
		return result
	}

	var nextRow []int
	for i := 0; i < len(matrixRow)-1; i++ {
		nextRow = append(nextRow, matrixRow[i+1]-matrixRow[i])
	}

	return append(result, makePredictionMatrix(nextRow)...)
}

func Solution2(filepath string) int {
	return 0
}

func parseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with: %v\n", numStr, err.Error())
	}

	return int(num)
}
