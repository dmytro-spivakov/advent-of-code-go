package main

import (
	aoc2023day01 "advent-of-code/2023/day01"
	aoc2023day02 "advent-of-code/2023/day02"
	aoc2023day03 "advent-of-code/2023/day03"
	aoc2023day04 "advent-of-code/2023/day04"
	aoc2023day05 "advent-of-code/2023/day05"
	aoc2023day06 "advent-of-code/2023/day06"
	aoc2023day07 "advent-of-code/2023/day07"
	aoc2023day08 "advent-of-code/2023/day08"
	aoc2023day09 "advent-of-code/2023/day09"
	aoc2023day10 "advent-of-code/2023/day10"
	aoc2023day11 "advent-of-code/2023/day11"
	aoc2023day12 "advent-of-code/2023/day12"
	aoc2023day13 "advent-of-code/2023/day13"
	aoc2023day14 "advent-of-code/2023/day14"
	aoc2023day15 "advent-of-code/2023/day15"
	aoc2023day16 "advent-of-code/2023/day16"
	aoc2023day17 "advent-of-code/2023/day17"
	aoc2023day18 "advent-of-code/2023/day18"
	aoc2023day19 "advent-of-code/2023/day19"
	aoc2024day01 "advent-of-code/2024/day01"
	"fmt"
	"time"
)

func runSolution(year, day, part int, s func() int) {
	start := time.Now()
	res := s()
	elapsed := time.Since(start)
	fmt.Printf("%d: Day %02d, solution %d: %v - %v\n", year, day, part, res, elapsed.Round(time.Microsecond))
}

func main() {
	// 2023
	runSolution(2023, 1, 1, func() int { return aoc2023day01.Solution1("2023/inputs/day01") })
	runSolution(2023, 1, 2, func() int { return aoc2023day01.Solution2("2023/inputs/day01") })

	runSolution(2023, 2, 1, func() int { return aoc2023day02.Solution1("2023/inputs/day02") })
	runSolution(2023, 2, 2, func() int { return aoc2023day02.Solution2("2023/inputs/day02") })

	runSolution(2023, 3, 1, func() int { return aoc2023day03.Solution1("2023/inputs/day03") })
	runSolution(2023, 3, 2, func() int { return aoc2023day03.Solution2("2023/inputs/day03") })

	runSolution(2023, 4, 1, func() int { return aoc2023day04.Solution1("2023/inputs/day04") })
	runSolution(2023, 4, 2, func() int { return aoc2023day04.Solution2("2023/inputs/day04") })

	runSolution(2023, 5, 1, func() int { return aoc2023day05.Solution1("2023/inputs/day05") })
	runSolution(2023, 5, 2, func() int { return aoc2023day05.Solution2("2023/inputs/day05") })

	runSolution(2023, 6, 1, func() int { return aoc2023day06.Solution1("2023/inputs/day06") })
	runSolution(2023, 6, 2, func() int { return aoc2023day06.Solution2("2023/inputs/day06") })

	runSolution(2023, 7, 1, func() int { return aoc2023day07.Solution1("2023/inputs/day07") })
	runSolution(2023, 7, 2, func() int { return aoc2023day07.Solution2("2023/inputs/day07") })

	runSolution(2023, 8, 1, func() int { return aoc2023day08.Solution1("2023/inputs/day08") })
	runSolution(2023, 8, 2, func() int { return aoc2023day08.Solution2("2023/inputs/day08") })

	runSolution(2023, 9, 1, func() int { return aoc2023day09.Solution1("2023/inputs/day09") })
	runSolution(2023, 9, 2, func() int { return aoc2023day09.Solution2("2023/inputs/day09") })

	runSolution(2023, 10, 1, func() int { return aoc2023day10.Solution1("2023/inputs/day10") })
	runSolution(2023, 10, 2, func() int { return aoc2023day10.Solution2("2023/inputs/day10") })

	runSolution(2023, 11, 1, func() int { return aoc2023day11.Solution1("2023/inputs/day11") })
	runSolution(2023, 11, 2, func() int { return aoc2023day11.Solution2("2023/inputs/day11") })

	runSolution(2023, 12, 1, func() int { return aoc2023day12.Solution1("2023/inputs/day12") })
	runSolution(2023, 12, 2, func() int { return aoc2023day12.Solution2("2023/inputs/day12") })

	runSolution(2023, 13, 1, func() int { return aoc2023day13.Solution1("2023/inputs/day13") })
	runSolution(2023, 13, 2, func() int { return aoc2023day13.Solution2("2023/inputs/day13") })

	runSolution(2023, 14, 1, func() int { return aoc2023day14.Solution1("2023/inputs/day14") })
	runSolution(2023, 14, 2, func() int { return aoc2023day14.Solution2("2023/inputs/day14") })

	runSolution(2023, 15, 1, func() int { return aoc2023day15.Solution1("2023/inputs/day15") })
	runSolution(2023, 15, 2, func() int { return aoc2023day15.Solution2("2023/inputs/day15") })

	runSolution(2023, 16, 1, func() int { return aoc2023day16.Solution1("2023/inputs/day16") })
	runSolution(2023, 16, 2, func() int { return aoc2023day16.Solution2("2023/inputs/day16") })

	runSolution(2023, 17, 1, func() int { return aoc2023day17.Solution1("2023/inputs/day17") })
	runSolution(2023, 17, 2, func() int { return aoc2023day17.Solution2("2023/inputs/day17") })

	runSolution(2023, 18, 1, func() int { return aoc2023day18.Solution1("2023/inputs/day18") })
	runSolution(2023, 18, 2, func() int { return aoc2023day18.Solution2("2023/inputs/day18") })

	runSolution(2023, 19, 1, func() int { return aoc2023day19.Solution1("2023/inputs/day19") })
	runSolution(2023, 19, 2, func() int { return aoc2023day19.Solution2("2023/inputs/day19") })

	// 2024
	runSolution(2024, 1, 1, func() int { return aoc2024day01.Solution1("2024/inputs/day01") })
	runSolution(2024, 1, 2, func() int { return aoc2024day01.Solution2("2024/inputs/day01") })
}
