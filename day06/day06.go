package day06

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time     int
	distance int
}

func (r Race) calcNewRecordOptions() int {
	var minWindUp, maxWindUp int

	for windUpTime := 0; windUpTime <= r.time; windUpTime++ {
		if distance := r.calcDistance(windUpTime); distance > r.distance {
			minWindUp = windUpTime
			break
		}
	}

	for windUpTime := r.time; windUpTime >= 0; windUpTime-- {
		if distance := r.calcDistance(windUpTime); distance > r.distance {
			maxWindUp = windUpTime
			break
		}
	}

	return maxWindUp - minWindUp + 1 // +1 to make the range inclusive on both ends, e.g min = 2, max = 4 -> 2,3,4 - 3 options
}

func (r Race) calcDistance(windUpTime int) int {
	speed := windUpTime // intial 0 + 1 per windUpTime unit
	distance := (r.time - windUpTime) * speed
	if distance < 0 {
		log.Fatalln("bruh")
	}

	return distance
}

func Solution1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	raceTimes := parseInts(scanner.Text())
	scanner.Scan()
	raceDistrances := parseInts(scanner.Text())
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err)
	}
	if timesLen, distancesLen := len(raceTimes), len(raceDistrances); timesLen == 0 || timesLen != distancesLen {
		log.Fatalf("Error during input file read/parsing - malformed return. Time entries: %d, distance entries: %d\n", timesLen, distancesLen)
	}

	var races []Race
	for i, time := range raceTimes {
		races = append(races, Race{time: time, distance: raceDistrances[i]})
	}

	result := 1
	for _, race := range races {
		result *= race.calcNewRecordOptions()
	}
	return result
}

func Solution2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %v\n", err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	raceTime := parseScatteredInt(scanner.Text())
	scanner.Scan()
	raceDistrance := parseScatteredInt(scanner.Text())
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err)
	}

	race := Race{time: raceTime, distance: raceDistrance}

	return race.calcNewRecordOptions()
}

func parseInts(numbersStr string) []int {
	var result []int

	numberRegex := regexp.MustCompile(`[0-9]+`)
	matches := numberRegex.FindAllString(numbersStr, -1)
	for _, match := range matches {
		result = append(result, parseInt(match))
	}

	return result
}

func parseScatteredInt(scatteredNumberStr string) int {
	number := 0
	numberRegex := regexp.MustCompile(`[0-9]+`)

	matches := numberRegex.FindAllString(scatteredNumberStr, -1)
	for _, match := range matches {
		number *= int(math.Pow10(len(match)))
		number += parseInt(match)
	}

	return number
}

func parseInt(numberStr string) int {
	result, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with error: %v\n", numberStr, err.Error())
	}

	return int(result)
}
