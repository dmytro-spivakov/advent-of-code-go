package day01

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solution1(filepath string) int {
	l, r := readInput(filepath)
	slices.Sort(l)
	slices.Sort(r)

	result := 0
	similarMap := make(map[int]int)
	for i := 0; i < len(l); i++ {
		diff := r[i] - l[i]
		if diff < 0 {
			diff *= -1
		}

		similarMap[r[i]] += 1
		result += diff
	}

	return result
}

func Solution2(filepath string) int {
	l, r := readInput(filepath)
	slices.Sort(l)
	slices.Sort(r)

	similarMap := make(map[int]int)
	for i := 0; i < len(l); i++ {
		similarMap[r[i]] += 1
	}

	result := 0
	for _, n := range l {
		result += n * similarMap[n]
	}

	return result
}

func readInput(filepath string) (left, right []int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %s\n", err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputNumPair := strings.SplitN(scanner.Text(), " ", 2)
		if len(inputNumPair) != 2 {
			log.Fatalf("Failed to parse input numbers pair %s\n", scanner.Text())
		}
		l, r := parseInt(strings.TrimSpace(inputNumPair[0])), parseInt(strings.TrimSpace(inputNumPair[1]))

		left = append(left, l)
		right = append(right, r)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %s\n", err.Error())
	}

	if len(left) != len(right) {
		log.Fatal("left and right length dont match")
	}

	return left, right
}

func parseInt(strNum string) int {
	num, err := strconv.ParseInt(strNum, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %s with: %s\n", err.Error())
	}

	return int(num)
}
