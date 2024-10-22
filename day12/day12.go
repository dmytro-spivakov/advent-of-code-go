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

var solution1ChunkCache = make(map[int][]string)

func Solution1Alt(filepath string) int {
	rows := parseInput(filepath)

	result := 0
	for _, r := range rows {
		result += countCombs(r.rawSeq, r.damagedGroups)
	}
	return result
}

/*
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
*/
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

func Solution1(filepath string) int {
	rows := parseInput(filepath)

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
		// chunkVars[idx] = getPossibleCombinations(length, "")
		chunkVars[idx] = getPossibleCombinations(length)
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

func getPossibleCombinations(length int) []string {
	if length < 0 {
		log.Fatalln("bruh")
	}

	if cached := solution1ChunkCache[length]; cached != nil {
		return cached
	}

	if length == 1 {
		return []string{".", "#"}
	}

	var combs []string

	for _, opt := range getPossibleCombinations(length - 1) {
		combs = append(combs, "."+opt)
		combs = append(combs, "#"+opt)
	}

	if solution1ChunkCache[length] == nil {
		solution1ChunkCache[length] = combs
	}
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
