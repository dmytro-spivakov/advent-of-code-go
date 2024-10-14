package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		log.Fatalf("Failed to open the input file with %v\n", err)
	}

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			log.Fatalf("Input file read failed with %v", err)
		}

		currentLine := scanner.Text()

		regExp := regexp.MustCompile(`\d{1}`)
		digitStrings := regExp.FindAllString(currentLine, -1)

		numberString := digitStrings[0] + digitStrings[len(digitStrings)-1]

		result += parseNumber(numberString)
	}

	return result
}

func parseNumber(s string) int {
	digit, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse string as Int: %v", s)
	}

	return int(digit)
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		log.Fatalf("Failed to open the input file with %v\n", err)
	}

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			log.Fatalf("Input file read failed with %v", err)
		}

		currentLine := scanner.Text()

		lIndex := len(currentLine) + 1
		rIndex := -1
		var lDigit, rDigit string

		for k, v := range getNumberStringToDigitMap() {
			if currentLIndex := strings.Index(currentLine, k); currentLIndex >= 0 && currentLIndex < lIndex {
				lIndex = currentLIndex
				lDigit = v
			}

			if currentRIndex := strings.LastIndex(currentLine, k); currentRIndex >= 0 && currentRIndex > rIndex {
				rIndex = currentRIndex
				rDigit = v
			}
		}

		result += parseNumber(lDigit + rDigit)
	}

	return result
}

func getNumberStringToDigitMap() map[string]string {
	numberNameDigitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for i := 0; i < 10; i++ {
		digitString := fmt.Sprint(i)
		numberNameDigitMap[digitString] = digitString
	}

	return numberNameDigitMap
}
