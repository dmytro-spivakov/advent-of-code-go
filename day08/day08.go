package day08

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Instruction struct {
	sequence    []string
	currentStep int
}

func makeInstruction(sequenceStr string) Instruction {
	sequenceStr = strings.TrimSpace(sequenceStr)
	return Instruction{
		sequence:    strings.Split(sequenceStr, ""),
		currentStep: 0,
	}
}

func (i *Instruction) next() string {
	index := i.currentStep % len(i.sequence)
	i.currentStep += 1
	return i.sequence[index]
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instruction := makeInstruction(scanner.Text())
	scanner.Scan()

	adjacentMap := make(map[string][2]string)
	for scanner.Scan() {
		adjacentRegex := regexp.MustCompile(`(\w{3}) = [(](\w{3}), (\w{3})[)]`)
		matches := adjacentRegex.FindStringSubmatch(scanner.Text())
		key := matches[1]
		values := [2]string{matches[2], matches[3]}

		adjacentMap[key] = values
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err)
	}

	iterationLimit := 10_000_000
	currentStep := "AAA"
	for i := 0; i < iterationLimit; i++ {
		if currentStep == "ZZZ" {
			return instruction.currentStep
		}

		currentInstruction := instruction.next()
		switch currentInstruction {
		case "L":
			currentStep = adjacentMap[currentStep][0]
		case "R":
			currentStep = adjacentMap[currentStep][1]
		default:
			log.Fatalf("Unknown instruction %v\n", currentInstruction)
		}
	}

	return -1
}

func Solution2(filepath string) int {
	return 0
}
