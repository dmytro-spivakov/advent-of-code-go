package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	printRows(rows)

	result := 0
	for _, r := range rows {
		result += countCombinations(r)
	}
	return result
}

func Solution2(filepath string) int {
	return -1
}

func countCombinations(r Row) int {
	unsolvedChunks := getUnsolvedChunkRanges(r.rawSeq)

	// { startIdx: []{all possible substrings of len()=length}, ... }
	chunkVars := make(map[int][]string)
	for idx, length := range unsolvedChunks {
		chunkVars[idx] = getPossibleCombinations(length, "")
	}

	// chunkVars to generate full row strings of all potential combinations of . and # in place of ?
	var chunksOrder []int
	for chunkIdx, _ := range chunkVars {
		chunksOrder = append(chunksOrder, chunkIdx)
	}
	slices.Sort(chunksOrder)

	varStrings := getAllPossibleStrings(strings.Join(r.rawSeq, ""), chunkVars, chunksOrder)
	// fmt.Println("ALL STRINGS:")
	// for _, vs := range varStrings {
	// 	fmt.Println(vs)
	// }
	// fmt.Println("--------------")

	// count valid strings
	result := 0
	expectedGroups := r.damagedGroups

	for _, v := range varStrings {
		dmgRegex := regexp.MustCompile(`#+`)
		matches := dmgRegex.FindAllString(v, -1)
		var groups []int
		for _, m := range matches {
			groups = append(groups, len(m))
		}

		if slices.Equal(expectedGroups, groups) {
			result += 1
		}
	}

	return result
}

func getUnsolvedChunkRanges(s []string) map[int]int {
	// { startIdx: length }
	unsolvedChunks := make(map[int]int)

	start := -1
	currentSlice := false
	for i, char := range s {
		if char == "?" {
			if !currentSlice {
				currentSlice = true
				start = i
			}
		} else {
			if currentSlice {
				unsolvedChunks[start] = i - start
				currentSlice = false
				start = -1
			}
		}
	}
	if currentSlice {
		unsolvedChunks[start] = len(s) - start
	}

	return unsolvedChunks
}

func getAllPossibleStrings(baseString string, chunkVariations map[int][]string, chunksOrder []int) []string {
	var result []string

	if len(chunksOrder) == 0 {
		return []string{baseString}
	}

	idx := chunksOrder[0]
	for _, v := range chunkVariations[idx] {
		newBaseString := ""
		if idx > 0 {
			newBaseString = baseString[:idx]
		}
		newBaseString += v
		if insertEndIdx := idx + len([]rune(v)); insertEndIdx < len([]rune(baseString)) {
			newBaseString += baseString[insertEndIdx:]
		}

		var newOrderChunk []int
		if len(chunksOrder) > 0 {
			newOrderChunk = chunksOrder[1:]
		} else {
			newOrderChunk = make([]int, 0)
		}
		result = append(
			result,
			getAllPossibleStrings(newBaseString, chunkVariations, newOrderChunk)...,
		)
	}
	return result
}

func getPossibleCombinations(length int, currentComb string) []string {
	var combs []string

	if length < 0 {
		log.Fatalln("bruh")
	}

	if length == 1 {
		combA := currentComb + "."
		combB := currentComb + "#"
		return []string{combA, combB}
	}

	combs = append(combs, getPossibleCombinations(length-1, currentComb+".")...)
	combs = append(combs, getPossibleCombinations(length-1, currentComb+"#")...)

	return combs
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
