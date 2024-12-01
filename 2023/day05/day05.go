package day05

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MappingRange struct {
	srcStart  int
	destStart int
	length    int
}

func (mr MappingRange) srcEnd() int {
	return mr.srcStart + mr.length - 1
}

func (mr MappingRange) destEnd() int {
	return mr.destStart + mr.length - 1
}

func (mr MappingRange) findDest(src int) int {
	if src < mr.srcStart {
		return -1
	}

	diff := src - mr.srcStart
	if diff > mr.length {
		return -1
	}

	return mr.destStart + diff
}

func (mr MappingRange) describe() string {
	return fmt.Sprintf("MappingRange { srcStart: %d, srcEnd: %d, destStart: %d, destEnd: %d }\n", mr.srcStart, mr.srcEnd(), mr.destStart, mr.destEnd())
}

type InputRange struct {
	start  int
	length int
}

func (ir InputRange) end() int {
	return ir.start + ir.length - 1
}

func (ir InputRange) describe() string {
	return fmt.Sprintf("InputRange { start: %d, end: %d }\n", ir.start, ir.end())
}

func (ir InputRange) applyMapping(mr MappingRange) (mappedInputs []InputRange, unmappedInputs []InputRange, ok bool) {
	// no overlap
	if ir.start > mr.srcEnd() || ir.end() < mr.srcStart {
		return mappedInputs, []InputRange{ir}, false
	}

	startDiff := ir.start - mr.srcStart
	endDiff := ir.end() - mr.srcEnd()

	// cut out the left outer range from ir and append it to the result
	if startDiff < 0 {
		startDiff = int(math.Abs(float64(startDiff)))
		leftOuterRange := InputRange{start: ir.start, length: startDiff}
		unmappedInputs = append(unmappedInputs, leftOuterRange)

		ir.start = mr.srcStart
		ir.length -= startDiff
	}

	// cut out the right outer range from ir and append it to the result
	if endDiff > 0 {
		rightOuterRange := InputRange{start: mr.srcEnd() + 1, length: endDiff}
		unmappedInputs = append(unmappedInputs, rightOuterRange)

		ir.length -= endDiff
	}

	// what's left of ir is contained in the mapping range
	// new ir start is mr.destStart + start offset
	startDiff = ir.start - mr.srcStart
	if startDiff < 0 {
		log.Fatalln("InputRange.applyMapping() you fucked up")
	}
	ir.start = mr.destStart + startDiff
	mappedInputs = append(mappedInputs, ir)

	return mappedInputs, unmappedInputs, true
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file %v\n", filepath)
	}

	// read seeds, handle the first two lines as an edge-case outside the main reading loop
	var mappedInputs []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedRegex := regexp.MustCompile(`\d+`)
	for _, seed := range seedRegex.FindAllString(scanner.Text(), -1) {
		mappedInputs = append(mappedInputs, parseInt(seed))
	}
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatalf("Error during input file read: %v\n", scanner.Err().Error())
	}

	var currentMappingSection []string
	for scanner.Scan() {
		currentLine := scanner.Text()
		// sections are ordered -> don't need to care about parsing headers
		if strings.Contains(currentLine, "map:") {
			continue
		}

		if len(currentLine) == 0 {
			mappedInputs = parseAndApplySectionMapping(currentMappingSection, mappedInputs)
			currentMappingSection = nil

			continue
		}

		currentMappingSection = append(currentMappingSection, currentLine)
	}
	// last section edge-case, there's no trailing empty line to trigger mapping
	mappedInputs = parseAndApplySectionMapping(currentMappingSection, mappedInputs)

	result := mappedInputs[0]
	for _, mappedInput := range mappedInputs {
		if mappedInput < result {
			result = mappedInput
		}
	}
	return result
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file %v\n", filepath)
	}

	// read seeds, handle the first two lines as an edge-case outside the main reading loop
	var mappedInputs []InputRange
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	mappedInputs = parseSeeds(scanner.Text())
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatalf("Error during input file read: %v\n", scanner.Err().Error())
	}

	var currentMappingSection []string
	for scanner.Scan() {
		currentLine := scanner.Text()
		// sections are ordered -> don't need to care about parsing headers
		if strings.Contains(currentLine, "map:") {
			continue
		}

		if len(currentLine) == 0 {
			mappedInputs = parseAndApplyRangeSectionMapping(currentMappingSection, mappedInputs)
			currentMappingSection = nil

			continue
		}

		currentMappingSection = append(currentMappingSection, currentLine)
	}
	// last section edge-case, there's no trailing empty line to trigger mapping
	mappedInputs = parseAndApplyRangeSectionMapping(currentMappingSection, mappedInputs)

	result := mappedInputs[0].start
	for _, mappedInput := range mappedInputs {
		if start := mappedInput.start; start < result {
			result = start
		}
	}
	return result
}

func parseSeeds(seedsLine string) (seeds []InputRange) {
	seedRegex := regexp.MustCompile(`(\d+)\s+(\d+)`)
	for _, seedMatches := range seedRegex.FindAllStringSubmatch(seedsLine, -1) {
		start := parseInt(seedMatches[1])
		length := parseInt(seedMatches[2])
		seeds = append(seeds, InputRange{start: start, length: length})
	}

	return seeds
}

func parseAndApplySectionMapping(sectionLines []string, inputs []int) (mappedInputs []int) {
	rangeRegex := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)

	var mappingRanges []MappingRange
	for _, sectionLine := range sectionLines {
		matches := rangeRegex.FindStringSubmatch(sectionLine)

		mappingRanges = append(
			mappingRanges,
			MappingRange{
				destStart: parseInt(matches[1]),
				srcStart:  parseInt(matches[2]),
				length:    parseInt(matches[3]),
			},
		)
	}

	for _, input := range inputs {
		found := false
		for _, mappingRange := range mappingRanges {
			if mappedInput := mappingRange.findDest(input); mappedInput >= 0 {
				mappedInputs = append(mappedInputs, mappedInput)
				found = true
				break
			}
		}

		if !found {
			mappedInputs = append(mappedInputs, input)
		}
	}

	return mappedInputs
}

func parseAndApplyRangeSectionMapping(sectionLines []string, inputs []InputRange) (mappedInputs []InputRange) {
	rangeRegex := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)

	var mappingRanges []MappingRange
	for _, sectionLine := range sectionLines {
		matches := rangeRegex.FindStringSubmatch(sectionLine)

		mappingRanges = append(
			mappingRanges,
			MappingRange{
				destStart: parseInt(matches[1]),
				srcStart:  parseInt(matches[2]),
				length:    parseInt(matches[3]),
			},
		)
	}

	for _, input := range inputs {
		mappedInputs = append(mappedInputs, applyRangeSectionMapping(mappingRanges, input)...)
	}
	return mappedInputs
}

func applyRangeSectionMapping(mappingRanges []MappingRange, inputRange InputRange) (mappedInputs []InputRange) {
	var unmappedInputs []InputRange

	newUnmappedInputs := false
	for _, mappingRange := range mappingRanges {
		if mapped, unmapped, ok := inputRange.applyMapping(mappingRange); ok {
			mappedInputs = append(mappedInputs, mapped...)
			unmappedInputs = append(unmappedInputs, unmapped...)
			newUnmappedInputs = true
			break
		}
	}

	if newUnmappedInputs {
		for _, unmapped := range unmappedInputs {
			mappedInputs = append(mappedInputs, applyRangeSectionMapping(mappingRanges, unmapped)...)
		}
		return mappedInputs
	} else {
		mappedInputs = append(mappedInputs, inputRange)
		return mappedInputs
	}
}

func parseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse int %v with err: %v\n", numStr, err)
	}

	return int(num)
}
