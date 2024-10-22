package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Row struct {
	rawSeq        []string
	damagedGroups []int
}

func Solution1(filepath string) int {
	rows := parseInput(filepath)

	result := 0
	for _, r := range rows {
		result += countCombs(r.rawSeq, r.damagedGroups)
	}
	return result
}

func countCombs(s []string, g []int) int {
	// base case - end of the seq
	if len(s) == 0 {
		if len(g) == 0 {
			return 1
		} else {
			// reached the end of the string, unallocated groups remain -> invalid
			return 0
		}
	}

	// base case = no more groups to alloc
	if len(g) == 0 {
		if !slices.Contains(s, "#") {
			return 1
		} else {
			return 0
		}
	}

	// ? eq . or # branching:
	// ? eq .
	result := 0
	if c := s[0]; c == "." || c == "?" {
		result += countCombs(s[1:], g)
	}

	// ? eq #
	if c := s[0]; c == "#" || c == "?" {
		// conditions for the next iter:
		// - the current group fits in the remaining s
		// - the next groupSize chars are all # or ?
		// - the group is followed either by end of the string or anything but #
		if gSize := g[0]; len(s) >= gSize && !slices.Contains(s[0:gSize], ".") && (len(s) == gSize || s[gSize] != "#") {
			// damaged group + trailing .
			newSIdx := gSize + 1
			if newSIdx <= len(s)-1 {
				result += countCombs(s[gSize+1:], g[1:])
			} else {
				result += countCombs(make([]string, 0), g[1:])
			}
		}

	}

	return result
}

func Solution2(filepath string) int {
	return -1
}

func parseInput(filepath string) []Row {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var m []Row
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		divInput := strings.Split(scanner.Text(), " ")
		seq := strings.Split(divInput[0], "")
		dmgGroups := strings.Split(divInput[1], ",")
		m = append(m, Row{rawSeq: seq, damagedGroups: parseInts(dmgGroups)})
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return m
}

func printRows(rows []Row) {
	fmt.Println("-----MATRIX START-----")
	for i, row := range rows {
		fmt.Printf("Row %d: %v | %v\n", i, strings.Join(row.rawSeq, ""), row.damagedGroups)
	}
	fmt.Println("------MATRIX END-----")
}

func parseInts(sNums []string) []int {
	var nums []int
	for _, sNum := range sNums {
		num64, err := strconv.ParseInt(sNum, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse number %v with: %v\n", sNum, err.Error())
		}

		nums = append(nums, int(num64))
	}

	return nums
}
