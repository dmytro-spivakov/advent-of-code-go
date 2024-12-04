package day19

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

type Part struct {
	x                int
	m                int
	a                int
	s                int
	appliedWorkflows []string
	workflowStep     string
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

func (p *Part) isProcessed() bool {
	if ws := (*p).workflowStep; ws == "A" || ws == "R" {
		return true
	}
	return false
}

func (p *Part) rating() int {
	return p.x + p.m + p.a + p.s
}

type RuleSet struct {
	name          string
	rawInput      string
	workflowSteps []func(p Part) (bool, string)
}

func (rs *RuleSet) apply(p *Part) bool {
	for _, wfStep := range rs.workflowSteps {
		ok, nextStep := wfStep(*p)
		if ok && !slices.Contains(p.appliedWorkflows, nextStep) {
			fmt.Printf("PART: %v, NEXT STEP: %v\n", p, nextStep)
			(*p).appliedWorkflows = append((*p).appliedWorkflows, p.workflowStep)
			(*p).workflowStep = nextStep
			return true
		}

	}
	return false
}

func Solution1(filepath string) int {
	rss, p := parseInput(filepath)
	fmt.Printf("RuleSet: len=%d, val=%v\n", len(rss), rss)
	fmt.Printf("Part: len=%d, val=%v\n", len(p), p)

	result := 0
	for len(p) > 0 {
		fmt.Printf("LEN=%v\n", len(p))
		currentPart := &p[0]
		fmt.Printf("currentPart=%v\n", currentPart)
		for _, rs := range rss {
			ok := rs.apply(currentPart)
			if ok {
				break
			}
		}

		if currentPart.isProcessed() {
			// fmt.Printf("currentPart=%v is processed\n", currentPart)
			if currentPart.workflowStep == "A" {
				result += currentPart.rating()
			}
			curLen := len(p)
			p[0] = p[curLen-1]
			p = p[:curLen-1]
		}
	}
	return result
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

		if len(currentLine) == 0 {
			continue
		}

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
	name := strings.Split(rawInput, "{")[0]
	ruleSet := RuleSet{
		name:     name,
		rawInput: rawInput,
	}

	ruleSegRegex := regexp.MustCompile(`\{(.+)\}`)
	ruleSegStr := ruleSegRegex.FindStringSubmatch(rawInput)[1]
	ruleSegments := strings.Split(ruleSegStr, ",")
	var rules []func(p Part) (bool, string)
	for _, ruleSeg := range ruleSegments {
		// samples:
		// a<10:gd - conditinal transition
		// A       - unconditional transition
		if !strings.Contains(ruleSeg, ":") {
			rules = append(rules, func(p Part) (bool, string) { return true, ruleSeg })
			continue
		}

		segParts := strings.Split(ruleSeg, ":")
		if len(segParts) != 2 {
			log.Fatalf("Failed to parse segment rule %s\n", ruleSeg)
		}

		endState := segParts[1]

		condition := []rune(segParts[0])
		conditionField := string(condition[0])                // a, m, s, x
		greaterThan := condition[1] == '>'                    // always < or >
		conditionThreshold := parseInt(string(condition[2:])) // whatever is left is the number to compare against

		// fmt.Printf("ruleSeg=%v, endState=%v, conditionFiled=%v, greaterThan=%v, conditionThreshold=%v\n", ruleSeg, endState, conditionField, greaterThan, conditionThreshold)

		if greaterThan {
			rules = append(
				rules,
				func(p Part) (bool, string) { ok := p.getVal(conditionField) > conditionThreshold; return ok, endState },
			)
		} else {
			rules = append(
				rules,
				func(p Part) (bool, string) { ok := p.getVal(conditionField) < conditionThreshold; return ok, endState },
			)
		}
	}
	ruleSet.workflowSteps = rules
	return ruleSet
}

func parsePart(rawInput string) Part {
	numRegex := regexp.MustCompile(`\d+`)
	matches := numRegex.FindAllString(rawInput, -1)
	if len(matches) != 4 {
		log.Fatalf("Failed to parse Part properly: input=%v, matches=%v\n", rawInput, matches)
	}

	// fields are always specified in the same order in the input
	return Part{
		x:            parseInt(matches[0]),
		m:            parseInt(matches[1]),
		a:            parseInt(matches[2]),
		s:            parseInt(matches[3]),
		workflowStep: "in",
	}
}

func parseInt(strNum string) int {
	num, err := strconv.ParseInt(strNum, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with %v\n", strNum, err.Error())
	}

	return int(num)
}
