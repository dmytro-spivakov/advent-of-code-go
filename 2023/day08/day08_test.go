package day08

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 2,
		"test_input12": 6,
	}

	for inputFile, expectedResult := range cases {
		result := Solution1(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution1() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution1() = %d, OK\n", result)
		}
	}
}

func TestSolution2(t *testing.T) {
	cases := map[string]int{
		"test_input21": 6,
	}

	for inputFile, expectedResult := range cases {
		result := Solution2(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution2() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution2() = %d, OK\n", result)
		}
	}
}

func TestInstructionNext(t *testing.T) {
	cases := []string{
		"LLR",
		"RRLRLR",
		"LRRLRRRLLRRRLRRRLLRRLRRRLRRLRRLRLRLRLRLRLLRRRLRRLRLRRRLRRRLRLRRRLRLRRLRRRLRRRLRLLRRRLRLLLRLRRRLRLRRLRRLLLLRRLRRLRLRLRRLRLRRLRRRLRRRLRLRLRRLLLLRRLRLRRLLRRRLRLRLRLRRRLRLLLRLRLRRRLRLRRRLRRRLRRRLLRRLRRRLRRRLRRRLRRRLRLLRRRLRLRRRLRLRLRRRLRRLRRLLRRRLRRRLRRRLRLRLRLRRLRRRLRRLRLRLRLRRRR",
	}

	for _, testCase := range cases {
		instruction := makeInstruction(testCase)
		overflowsCount := 3

		result := ""
		for i := 0; i < overflowsCount*len(testCase); i++ {
			result += instruction.next()
		}

		expectedResult := ""
		for i := 0; i < overflowsCount; i++ {
			expectedResult += testCase
		}

		if result != expectedResult {
			t.Fatalf("Instruction.next() = %s, expecting %s\n", result, expectedResult)
		} else {
			fmt.Printf("Instruction.next() = %s, OK\n", result)
		}
	}

}
