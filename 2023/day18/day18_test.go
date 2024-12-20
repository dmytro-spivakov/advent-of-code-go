package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 62,
	}

	for input, expectedResult := range cases {
		result := Solution1(input)

		assert.Equal(t, expectedResult, result, "Solution1()")
	}
}

func TestSolution2(t *testing.T) {
	cases := map[string]int{
		"test_input11": 952408144115,
	}

	for input, expectedResult := range cases {
		result := Solution2(input)

		assert.Equal(t, expectedResult, result, "Solution2()")
	}
}
