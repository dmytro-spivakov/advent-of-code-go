package main

import (
	"advent-of-code/day01"
	"advent-of-code/day02"
	"advent-of-code/day03"
	"advent-of-code/day04"
	"advent-of-code/day05"
	"advent-of-code/day06"
	"fmt"
)

func main() {
	fmt.Printf("Day 1, solution 1: %v\n", day01.Solution1("inputs/day01"))
	fmt.Printf("Day 1, solution 2: %v\n", day01.Solution2("inputs/day01"))

	fmt.Printf("Day 2, solution 1: %v\n", day02.Solution1("inputs/day02"))
	fmt.Printf("Day 2, solution 2: %v\n", day02.Solution2("inputs/day02"))

	fmt.Printf("Day 3, solution 1: %v\n", day03.Solution1("inputs/day03"))
	fmt.Printf("Day 3, solution 2: %v\n", day03.Solution2("inputs/day03"))

	fmt.Printf("Day 4, solution 1: %v\n", day04.Solution1("inputs/day04"))
	fmt.Printf("Day 4, solution 2: %v\n", day04.Solution2("inputs/day04"))

	fmt.Printf("Day 5, solution 1: %v\n", day05.Solution1("inputs/day05"))
	fmt.Printf("Day 5, solution 2: %v\n", day05.Solution2("inputs/day05"))

	fmt.Printf("Day 6, solution 1: %v\n", day06.Solution1("inputs/day06"))
}
