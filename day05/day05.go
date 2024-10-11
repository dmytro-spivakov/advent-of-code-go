package day05

import (
	"bufio"
	"log"
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

type Range struct {
	srcStart  int
	destStart int
	length    int
}

func (r Range) findDest(src int) int {
	if src < r.srcStart {
		return -1
	}

	diff := src - r.srcStart
	if diff > r.length {
		return -1
	}

	return r.destStart + diff
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file %v\n", filepath)
	}

	processingOrder := []string{
		"soil",
		"fertilizer",
		"water",
		"light",
		"temperature",
		"humidity",
		"location",
	}

	processStepToHeader := map[string]string{
		"soil":        "seed-to-soil map:",
		"fertilizer":  "soil-to-fertilizer map:",
		"water":       "fertilizer-to-water map:",
		"light":       "water-to-light map:",
		"temperature": "light-to-temperature map:",
		"humidity":    "temperature-to-humidity map:",
		"location":    "humidity-to-location map:",
	}
	var seeds []int
	processStepToRange := make(map[string][]Range)

	rangeRegex := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
	seedRegex := regexp.MustCompile(`\d+`)
	currentContext := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		// sections are separated by a blank line
		if len(currentLine) == 0 {
			currentContext = ""
			continue
		}

		if strings.Contains(currentLine, "seeds:") {
			for _, seed := range seedRegex.FindAllString(currentLine, -1) {
				seeds = append(seeds, parseInt(seed))
			}
			continue
		}

		contextChanged := false
		for step, header := range processStepToHeader {
			if strings.Contains(currentLine, header) {
				currentContext = step
				contextChanged = true
				break
			}
		}
		if contextChanged {
			continue
		}

		rangeNumbers := rangeRegex.FindStringSubmatch(currentLine)
		if rangeNumCount := len(rangeNumbers); rangeNumCount > 0 && rangeNumCount != 4 {
			log.Fatalf("Failed to parse range row %v\n", rangeNumbers)
		}

		processStepToRange[currentContext] = append(
			processStepToRange[currentContext],
			Range{
				destStart: parseInt(rangeNumbers[1]),
				srcStart:  parseInt(rangeNumbers[2]),
				length:    parseInt(rangeNumbers[3]),
			},
		)
	}
	if scanner.Err() != nil {
		log.Fatalf("Error during input file read: %v\n", scanner.Err().Error())
	}

	mappingInput := seeds
	var outputs []int
	for _, step := range processingOrder {
		outputs = nil

		for _, input := range mappingInput {
			found := false
			for _, rangeStr := range processStepToRange[step] {
				if mappedValue := rangeStr.findDest(input); mappedValue >= 0 {
					outputs = append(outputs, mappedValue)
					found = true
					break
				}
			}

			if !found {
				outputs = append(outputs, input)
			}
		}

		mappingInput = make([]int, len(outputs))
		copy(mappingInput, outputs)
	}

	result := outputs[0]
	for _, output := range outputs {
		if output < result {
			result = output
		}
	}

	return result
}

func parseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse int %v with err: %v\n", numStr, err)
	}

	return int(num)
}
