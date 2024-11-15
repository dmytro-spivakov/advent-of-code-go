package day17

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Value [6]int
	index int // Used internally by the priority queue
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Item

// Basic heap operations
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Lower values[0] have higher priority
	return pq[i].Value[0] < pq[j].Value[0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the item with lowest Value[0]
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

// Enqueue adds a [6]int array to the queue
func (pq *PriorityQueue) Enqueue(value [6]int) {
	item := &Item{
		Value: value,
	}
	heap.Push(pq, item)
}

// Dequeue removes and returns the [6]int array with lowest value[0]
func (pq *PriorityQueue) Dequeue() [6]int {
	if len(*pq) == 0 {
		panic("Priority queue is empty")
	}
	item := heap.Pop(pq).(*Item)
	return item.Value
}

// Peek returns the [6]int array with lowest value[0] without removing it
func (pq *PriorityQueue) Peek() [6]int {
	if len(*pq) == 0 {
		panic("Priority queue is empty")
	}
	return (*pq)[0].Value
}

// IsEmpty returns true if the queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func Solution1(filepath string) int {
	m := readInput(filepath)

	pq := PriorityQueue{}
	seen := make(map[string]bool)
	pq.Enqueue([6]int{0, 0, 0, 0, 0, 0}) // heat loss, y, x, dY, dX, n of steps in straight line

	for pq.Len() > 0 {
		el := pq.Dequeue()
		heatLoss, y, x, dY, dX, n := el[0], el[1], el[2], el[3], el[4], el[5]

		if y == len(m)-1 && x == len(m[y])-1 {
			return heatLoss
		}

		seenKey := fmt.Sprintf("%d;%d;%d;%d;%d", y, x, dY, dX, n)
		if seen[seenKey] {
			continue
		}
		seen[seenKey] = true

		// keep going in the same direction
		if n < 3 && [2]int{dY, dX} != [2]int{0, 0} {
			newY, newX := y+dY, x+dX
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[0]) {
				// do nothing, I don't want to invert this condition
			} else {
				pq.Enqueue([6]int{heatLoss + m[newY][newX], newY, newX, dY, dX, n + 1})
			}
		}

		// explore all the other directions except for the current and its reverse
		for _, diffs := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newDY, newDX := diffs[0], diffs[1]
			newY, newX := y+newDY, x+newDX
			if newDir := [2]int{newDY, newDX}; newDir == [2]int{dY, dX} || newDir == [2]int{-dY, -dX} {
				continue
			}
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[newY]) {
				continue
			}
			pq.Enqueue([6]int{heatLoss + m[newY][newX], newY, newX, newDY, newDX, 1})
		}

	}
	return -1
}

func Solution2(filepath string) int {
	m := readInput(filepath)

	pq := PriorityQueue{}
	seen := make(map[string]bool)
	pq.Enqueue([6]int{0, 0, 0, 0, 0, 0}) // heat loss, y, x, dY, dX, n of steps in straight line

	for pq.Len() > 0 {
		el := pq.Dequeue()
		heatLoss, y, x, dY, dX, n := el[0], el[1], el[2], el[3], el[4], el[5]
		// fmt.Printf("DEBUG: hl=%d, y=%d, x=%d, dY=%d, dX=%d, n=%d\n", heatLoss, y, x, dY, dX, n)
		if y == len(m)-1 && x == len(m[y])-1 && n >= 4 {
			return heatLoss
		}

		seenKey := fmt.Sprintf("%d;%d;%d;%d;%d", y, x, dY, dX, n)
		if seen[seenKey] {
			continue
		}
		seen[seenKey] = true

		// keep going in the same direction
		if n < 10 && [2]int{dY, dX} != [2]int{0, 0} {
			newY, newX := y+dY, x+dX
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[0]) {
				// do nothing, I don't want to invert this condition
			} else {
				pq.Enqueue([6]int{heatLoss + m[newY][newX], newY, newX, dY, dX, n + 1})
			}
		}

		if n != 0 && n < 4 {
			continue
		}
		// explore all the other directions except for the current and its reverse
		for _, diffs := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newDY, newDX := diffs[0], diffs[1]
			newY, newX := y+newDY, x+newDX
			if newDir := [2]int{newDY, newDX}; newDir == [2]int{dY, dX} || newDir == [2]int{-dY, -dX} {
				continue
			}
			if newY < 0 || newY >= len(m) || newX < 0 || newX >= len(m[newY]) {
				continue
			}
			pq.Enqueue([6]int{heatLoss + m[newY][newX], newY, newX, newDY, newDX, 1})
		}

	}
	return -1
}

func readInput(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var matrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentRow := strings.Split(scanner.Text(), "")
		currentRowNums := make([]int, len(currentRow))
		for i, el := range currentRow {
			currentRowNums[i] = parseInt(el)
		}
		matrix = append(matrix, currentRowNums)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return matrix
}

func parseInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with: %s\n", s, err.Error())
	}

	return int(num)
}

func printMatrix(m [][]int) {
	fmt.Println("-----MATRIX START-----")
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println("------MATRIX END------")
}
