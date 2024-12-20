package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 19114,
	}

	for input, expectedResult := range cases {
		result := Solution1(input)

		assert.Equal(t, expectedResult, result, "Solution1()")
	}
}
