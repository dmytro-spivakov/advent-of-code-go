package day15

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var segments []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		segments = append(segments, strings.Split(scanner.Text(), ",")...)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	result := 0
	for _, segment := range segments {
		result += segmentHash(segment)
	}

	return result
}

func segmentHash(s string) int {
	current := 0

	for _, r := range s {
		asciiCode := int(r)
		current += asciiCode
		current *= 17
		current = current % 256
	}

	return current
}

func Solution2(filepath string) int {
	return -1
}
