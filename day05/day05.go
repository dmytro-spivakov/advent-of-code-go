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

/*
--- Day 5: If You Give A Seed A Fertilizer ---

You take the boat and find the gardener right where you were told he would be: managing a giant "garden" that looks more to you like a farm.

"A water source? Island Island is the water source!" You point out that Snow Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with! Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand soon; we only turned off the water a few days... weeks... oh no." His face sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot to check why we stopped getting more sand! There's a ferry leaving soon that is headed over in that direction - it's much faster than your boat. Could you please go check it out?"

You barely have time to agree to this request when he brings up another. "While you wait for the ferry, maybe you can help us with our food production problem. The latest Island Island Almanac just arrived and we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil to use with each kind of seed, what type of fertilizer to use with each kind of soil, what type of water to use with each kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on is identified with a number, but numbers are reused by each category - that is, soil 123 and fertilizer 123 aren't necessarily related to each other.

For example:
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4

The almanac starts by listing which seeds need to be planted: seeds 79, 14, 55, and 13.

The rest of the almanac contains a list of maps which describe how to convert numbers from a source category into numbers in a destination category. That is, the section that starts with seed-to-soil map: describes how to convert a seed number (the source) to a soil number (the destination). This lets the gardener and his team know which soil to use with which seeds, which water to use with which fertilizer, and so on.

Rather than list every source number and its corresponding destination number one by one, the maps describe entire ranges of numbers that can be converted. Each line within a map contains three numbers: the destination range start, the source range start, and the range length.

Consider again the example seed-to-soil map:

50 98 2
52 50 48
The first line has a destination range start of 50, a source range start of 98, and a range length of 2. This line means that the source range starts at 98 and contains two values: 98 and 99. The destination range is the same length, but it starts at 50, so its two values are 50 and 51. With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48 values: 50, 51, ..., 96, 97. This corresponds to a destination range starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed number 53 corresponds to soil number 55.

Any source numbers that aren't mapped correspond to the same destination number. So, seed number 10 corresponds to soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers looks like this:
seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51
With this map, you can look up the soil number required for each initial seed number:

Seed number 79 corresponds to soil number 81.
Seed number 14 corresponds to soil number 14.
Seed number 55 corresponds to soil number 57.
Seed number 13 corresponds to soil number 13.

The gardener and his team want to get started as soon as possible, so they'd like to know the closest location that needs a seed. Using these maps, find the lowest location number that corresponds to any of the initial seeds. To do this, you'll need to convert each seed number through other categories until you can find its corresponding location number. In this example, the corresponding types are:

Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.
So, the lowest location number in this example is 35.

What is the lowest location number that corresponds to any of the initial seed numbers?
*/

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
	fmt.Println(startDiff)
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
	// TODO: rm debug
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

/*
--- Part Two ---

Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13
This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest location number that corresponds to any of the initial seed numbers?
*/

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
	fmt.Print("RESULTS")
	for _, mappedInput := range mappedInputs {
		fmt.Print(mappedInput.describe())
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
		fmt.Println("NEW UNMAPPED INPUTS:")
		for _, unmapped := range unmappedInputs {
			fmt.Println(unmapped.describe())
		}
		fmt.Println("END NEW UNMAPPED INPUTS")
		for _, unmapped := range unmappedInputs {
			mappedInputs = append(mappedInputs, applyRangeSectionMapping(mappingRanges, unmapped)...)
		}
		return mappedInputs
	} else {
		mappedInputs = append(mappedInputs, inputRange)
		fmt.Println("NO NEW UNMAPPED INPUTS:")
		for _, mapped := range mappedInputs {
			fmt.Println(mapped.describe())
		}
		fmt.Println("NO END NEW UNMAPPED INPUTS")
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
