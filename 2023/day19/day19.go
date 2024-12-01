package day19

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

func (p *Part) getVal(field string) int {
	switch field {
	case "x":
		return (*p).x
	case "m":
		return (*p).m
	case "a":
		return (*p).a
	case "s":
		return (*p).s
	}

	log.Fatalf("Part.getVal() unknown field %s\n", field)
	return -1
}

type RuleSet struct {
	name          string
	rawInput      string
	workflowSteps []func()
}

func (rs *RuleSet) init() {
	// parse `rawInput` and populate `workflowSteps`
}

func Solution1(filepath string) int {
	return -1
}

func Solution2(filepath string) int {
	return -1
}

func parseInput(filepath string) ([]RuleSet, []Part) {
	var ruleSets []RuleSet
	var parts []Part

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if strings.HasPrefix(currentLine, "{") {
			parts = append(parts, parsePart(currentLine))
		} else {
			ruleSets = append(ruleSets, parseRuleSet(currentLine))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return ruleSets, parts
}

func parseRuleSet(rawInput string) RuleSet {
	// name := strings.Split(rawInput, "{")[0]
	// ruleSet := RuleSet{
	// 	name:     name,
	// 	rawInput: rawInput,
	// }
	// ruleSegments := strings.Split(rawInput, ",")
	// var rules []func()
	// for _, ruleSeg := range ruleSegments {
	// samples:
	// a<10:gd - conditinal transition
	// A       - unconditional transition

	// rule func signature func(p Part) (applied bool, transitTo string)
	// }

	// return ruleSet
	return RuleSet{}
}

func parsePart(rawInput string) Part {
	numRegex := regexp.MustCompile(`\d+`)
	matches := numRegex.FindAllString(rawInput, -1)
	if len(matches) != 4 {
		log.Fatalf("Failed to parse Part properly: input=%v, matches=%v\n", rawInput, matches)
	}

	// fields are always specified in the same order in the input
	return Part{
		x: parseInt(matches[0]),
		m: parseInt(matches[1]),
		a: parseInt(matches[2]),
		s: parseInt(matches[3]),
	}
}

func parseInt(strNum string) int {
	num, err := strconv.ParseInt(strNum, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with %v\n", strNum, err.Error())
	}

	return int(num)
}
