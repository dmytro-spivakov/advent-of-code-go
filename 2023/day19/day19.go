package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Workflow struct {
	name     string
	fallback string
	rules    []Rule
}

type Rule struct {
	field       string
	greaterThan bool
	threshold   int
	next        string
}

func Solution1(filepath string) int {
	parts, workflows := parseInput(filepath)
	fmt.Println("Parts:")
	for _, p := range parts {
		fmt.Println(p)
	}
	fmt.Println("-------------------")
	fmt.Println("Workflows:")
	for _, wf := range workflows {
		fmt.Println(wf)
	}
	fmt.Println("-------------------")

	wfMap := make(map[string]Workflow)
	for _, wf := range workflows {
		wfMap[wf.name] = wf
	}

	result := 0
	for _, p := range parts {
		var history []string

		currentStep := "in"
		for currentStep != "A" && currentStep != "R" {
			wf := wfMap[currentStep]

			pass := false
			var nextStep string
			for _, r := range wf.rules {
				if (r.greaterThan && p[r.field] > r.threshold) || (!r.greaterThan && p[r.field] < r.threshold) {
					pass = true
					nextStep = r.next
					break
				}
			}

			history = append(history, currentStep)

			if !pass {
				nextStep = wf.fallback
			}
			currentStep = nextStep
		}

		history = append(history, currentStep)
		fmt.Println(strings.Join(history, "->"))
		if currentStep == "A" {
			result += p["x"] + p["m"] + p["a"] + p["s"]
		}

	}

	return result
}

func Solution2(filepath string) int {
	return -1
}

func parseInput(filepath string) ([]map[string]int, []Workflow) {
	var parts []map[string]int
	var workflows []Workflow

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if len(currentLine) == 0 {
			continue
		}

		if strings.HasPrefix(currentLine, "{") {
			parts = append(parts, parsePart(currentLine))
		} else {
			workflows = append(workflows, parseWorkflow(currentLine))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return parts, workflows
}

func parseWorkflow(rawInput string) Workflow {
	wfParts := strings.Split(rawInput[:len(rawInput)-1], "{")
	name := wfParts[0]
	rest := strings.Split(wfParts[1], ",")
	workflow := Workflow{
		name:     name,
		fallback: rest[len(rest)-1],
	}
	rest = rest[:len(rest)-1]
	for _, ruleStr := range rest {
		parts := strings.Split(ruleStr, ":")
		next := parts[1]
		greaterThan := strings.Contains(parts[0], ">")
		var field string
		var threshold int
		if greaterThan {
			split := strings.Split(parts[0], ">")
			field = split[0]
			threshold = parseInt(split[1])
		} else {
			split := strings.Split(parts[0], "<")
			field = split[0]
			threshold = parseInt(split[1])
		}
		workflow.rules = append(
			workflow.rules,
			Rule{
				field:       field,
				greaterThan: greaterThan,
				threshold:   threshold,
				next:        next,
			},
		)
	}

	return workflow
}

func parsePart(rawInput string) map[string]int {
	numRegex := regexp.MustCompile(`\d+`)
	matches := numRegex.FindAllString(rawInput, -1)
	if len(matches) != 4 {
		log.Fatalf("Failed to parse Part properly: input=%v, matches=%v\n", rawInput, matches)
	}

	// fields are always specified in the same order in the input
	return map[string]int{
		"x": parseInt(matches[0]),
		"m": parseInt(matches[1]),
		"a": parseInt(matches[2]),
		"s": parseInt(matches[3]),
	}
}

func parseInt(strNum string) int {
	num, err := strconv.ParseInt(strNum, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with %v\n", strNum, err.Error())
	}

	return int(num)
}
