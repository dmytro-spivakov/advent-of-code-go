package main

import (
	"advent-of-code/day01"
	"advent-of-code/day02"
	"advent-of-code/day03"
	"advent-of-code/day04"
	"advent-of-code/day05"
	"advent-of-code/day06"
	"advent-of-code/day07"
	"advent-of-code/day08"
	"advent-of-code/day09"
	"advent-of-code/day10"
	"advent-of-code/day11"
	"advent-of-code/day12"
	"fmt"
	"time"
)

func runSolution(day, part int, s func() int) {
	start := time.Now()
	res := s()
	elapsed := time.Since(start)
	fmt.Printf("Day %d, solution %d: %v - %v\n", day, part, res, elapsed.Round(time.Microsecond))
}

func main() {
	runSolution(1, 1, func() int { return day01.Solution1("inputs/day01") })
	runSolution(1, 2, func() int { return day01.Solution2("inputs/day01") })

	runSolution(2, 1, func() int { return day02.Solution1("inputs/day02") })
	runSolution(2, 2, func() int { return day02.Solution2("inputs/day02") })

	runSolution(3, 1, func() int { return day03.Solution1("inputs/day03") })
	runSolution(3, 2, func() int { return day03.Solution2("inputs/day03") })

	runSolution(4, 1, func() int { return day04.Solution1("inputs/day04") })
	runSolution(4, 2, func() int { return day04.Solution2("inputs/day04") })

	runSolution(5, 1, func() int { return day05.Solution1("inputs/day05") })
	runSolution(5, 2, func() int { return day05.Solution2("inputs/day05") })

	runSolution(6, 1, func() int { return day06.Solution1("inputs/day06") })
	runSolution(6, 2, func() int { return day06.Solution2("inputs/day06") })

	runSolution(7, 1, func() int { return day07.Solution1("inputs/day07") })
	runSolution(7, 2, func() int { return day07.Solution2("inputs/day07") })

	runSolution(8, 1, func() int { return day08.Solution1("inputs/day08") })
	runSolution(8, 2, func() int { return day08.Solution2("inputs/day08") })

	runSolution(9, 1, func() int { return day09.Solution1("inputs/day09") })
	runSolution(9, 2, func() int { return day09.Solution2("inputs/day09") })

	runSolution(10, 1, func() int { return day10.Solution1("inputs/day10") })
	runSolution(10, 2, func() int { return day10.Solution2("inputs/day10") })

	runSolution(11, 1, func() int { return day11.Solution1("inputs/day11") })
	runSolution(11, 2, func() int { return day11.Solution2("inputs/day11") })

	runSolution(12, 1, func() int { return day12.Solution1("inputs/day12") })
	runSolution(12, 2, func() int { return day12.Solution2("inputs/day12") })
}
