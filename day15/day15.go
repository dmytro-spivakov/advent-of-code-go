package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solution1(filepath string) int {
	segments := readInput(filepath)

	result := 0
	for _, segment := range segments {
		result += segmentHash(segment)
	}

	return result
}

func Solution2(filepath string) int {
	segments := readInput(filepath)

	boxes := make(map[int][]string)
	for _, segment := range segments {
		if strings.Contains(segment, "=") {
			label, focal := splitIntoTuple(segment, "=")
			fullLabel := fmt.Sprintf("%s %s", label, focal)

			boxIdx := segmentHash(label)
			replaceIdx := slices.IndexFunc(boxes[boxIdx], func(lense string) bool {
				return strings.HasPrefix(lense, label)
			})

			if replaceIdx >= 0 {
				boxes[boxIdx] = slices.Replace(boxes[boxIdx], replaceIdx, replaceIdx+1, fullLabel)
			} else {
				boxes[boxIdx] = append(boxes[boxIdx], fullLabel)
			}
		} else {
			label := segment[:len(segment)-1]
			boxIdx := segmentHash(label)

			removeIdx := slices.IndexFunc(boxes[boxIdx], func(lense string) bool {
				return strings.HasPrefix(lense, label)
			})
			if removeIdx < 0 {
				continue
			}

			boxes[boxIdx] = slices.Delete(boxes[boxIdx], removeIdx, removeIdx+1)
		}
	}

	result := 0
	for k, v := range boxes {
		boxNumMultiplier := k + 1
		for i, lense := range v {
			_, focal := splitIntoTuple(lense, " ")
			lenseSlot := i + 1
			result += boxNumMultiplier * lenseSlot * parseInt(focal)
		}
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

func splitIntoTuple(s, sep string) (string, string) {
	results := strings.SplitN(s, sep, 2)
	return results[0], results[1]
}

func parseInt(num string) int {
	res, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with: %v\n", num, err.Error())
	}

	return int(res)
}

func readInput(filepath string) []string {
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

	return segments
}
