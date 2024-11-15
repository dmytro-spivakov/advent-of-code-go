package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 102,
	}

	for input, expectedResult := range cases {
		result := Solution1(input)

		assert.Equal(t, expectedResult, result, "Solution1()")
	}
}

func TestSolution2(t *testing.T) {
	cases := map[string]int{
		"test_input11": 94,
		"test_input12": 71,
	}

	for input, expectedResult := range cases {
		result := Solution2(input)

		assert.Equal(t, expectedResult, result, "Solution2()")
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQueue{}
	pq.Enqueue([6]int{0, 0, 0, 0, 0, 0})
	pq.Enqueue([6]int{4, 0, 0, 0, 0, 0})
	pq.Enqueue([6]int{2, 0, 0, 0, 0, 0})
	pq.Enqueue([6]int{6, 0, 0, 0, 0, 0})
	pq.Enqueue([6]int{99, 0, 0, 0, 0, 0})
	pq.Enqueue([6]int{1, 0, 0, 0, 0, 0})

	cases := map[int]int{
		0: 0,
		1: 1,
		2: 2,
		3: 4,
		4: 6,
		5: 99,
	}

	for i := 0; i <= 5; i++ {
		heatLoss := pq.Dequeue()[0]
		assert.Equal(t, cases[i], heatLoss, "PQ Enqueue(), Dequeue()")
	}
	assert.Equal(t, 0, pq.Len(), "PQ Len()")
}
