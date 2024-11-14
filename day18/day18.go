package day18

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction  [2]int
	length     int
	direction2 [2]int
	length2    int
}

// func Solution1(filepath string) int {
// 	instructions := readInput(filepath)
// 	printInstructions(instructions)
//
// 	var m [][]string
// 	y, x := 0, 0
// 	m = append(m, []string{"#"})
//
// 	// draw 2d matrix following the trench instructions
// 	for _, instruction := range instructions {
// 		for i := 0; i < instruction.length; i++ {
// 			newY, newX := y+instruction.direction[0], x+instruction.direction[1]
//
// 			if newY < 0 {
// 				newTopRow := make([]string, len(m[0]))
// 				for i := 0; i < len(newTopRow); i++ {
// 					newTopRow[i] = "."
// 				}
// 				m = append([][]string{newTopRow}, m...)
// 				newY = 0
// 			}
// 			if newX < 0 {
// 				for y1 := 0; y1 < len(m); y1++ {
// 					m[y1] = append([]string{"."}, m[y1]...)
// 				}
// 				newX = 0
// 			}
// 			if newY >= len(m) {
// 				newBotRow := make([]string, len(m[0]))
// 				for i := 0; i < len(newBotRow); i++ {
// 					newBotRow[i] = "."
// 				}
// 				m = append(m, newBotRow)
// 			}
// 			if newX >= len(m[0]) {
// 				for y1 := 0; y1 < len(m); y1++ {
// 					m[y1] = append(m[y1], ".")
// 				}
// 			}
//
// 			y, x = newY, newX
// 			m[y][x] = "#"
// 		}
// 	}
// 	printMatrix(m)
//
// 	// fill the area surrounded by the trench with #
// 	for y := 0; y < len(m); y++ {
// 		isInside := false
// 		previousTileIsHash := false
// 		previousCorner := ""
// 		for x := 0; x < len(m[y]); x++ {
// 			if m[y][x] == "#" {
// 				if !previousTileIsHash {
// 					if y > 0 && y < len(m)-1 && m[y-1][x] == "#" && m[y+1][x] == "#" {
// 						isInside = !isInside
// 						previousCorner = "|"
// 					} else if y > 0 && m[y-1][x] == "#" {
// 						previousCorner = "L"
// 					} else if y < len(m)-1 && m[y+1][x] == "#" {
// 						previousCorner = "F"
// 					}
// 					previousTileIsHash = true
// 				} else {
// 					if y < len(m)-1 && previousCorner == "L" && m[y+1][x] == "#" {
// 						isInside = !isInside
// 						previousCorner = ""
// 					} else if y > 0 && previousCorner == "F" && m[y-1][x] == "#" {
// 						isInside = !isInside
// 						previousCorner = ""
// 					}
// 				}
//
// 			} else {
// 				previousCorner = ""
// 				previousTileIsHash = false
// 				if isInside {
// 					m[y][x] = "x"
// 				}
// 			}
// 		}
// 	}
//
// 	printMatrix(m)
//
// 	area := 0
// 	for y := 0; y < len(m); y++ {
// 		for x := 0; x < len(m[y]); x++ {
// 			if char := m[y][x]; char == "#" || char == "x" {
// 				area++
// 			}
// 		}
// 	}
// 	return area
// }

func Solution1(filepath string) int {
	instructions := readInput(filepath)

	points := [][2]int{{0, 0}}
	boundaryPoints := 0
	for _, inst := range instructions {
		dY, dX := inst.direction[0], inst.direction[1]
		scalar := inst.length
		boundaryPoints += scalar

		lastPoint := points[len(points)-1]
		lastY, lastX := lastPoint[0], lastPoint[1]

		points = append(points, [2]int{
			lastY + dY*scalar,
			lastX + dX*scalar,
		})
	}

	// magic
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		y1, x1 := points[i][0], points[i][1]
		y2, x2 := points[i+1][0], points[i+1][1]
		sum += (y1 + y2) * (x1 - x2)
	}

	return (sum / 2) - (boundaryPoints / 2) + 1 + boundaryPoints
}

func Solution2(filepath string) int {
	instructions := readInput(filepath)

	points := [][2]int{{0, 0}}
	boundaryPoints := 0
	for _, inst := range instructions {
		dY, dX := inst.direction2[0], inst.direction2[1]
		scalar := inst.length2
		boundaryPoints += scalar

		lastPoint := points[len(points)-1]
		lastY, lastX := lastPoint[0], lastPoint[1]

		points = append(points, [2]int{
			lastY + dY*scalar,
			lastX + dX*scalar,
		})
	}

	// magic
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		y1, x1 := points[i][0], points[i][1]
		y2, x2 := points[i+1][0], points[i+1][1]
		sum += (y1 + y2) * (x1 - x2)
	}

	return (sum / 2) - (boundaryPoints / 2) + 1 + boundaryPoints
}

func readInput(filepath string) []Instruction {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open the input file with: %v\n", err.Error())
	}

	var instructions []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// append
		currentLine := strings.Split(scanner.Text(), " ")
		directionName, length := currentLine[0], currentLine[1]
		directionNameToCoords := map[string][2]int{
			"U": {-1, 0},
			"D": {1, 0},
			"L": {0, -1},
			"R": {0, 1},
		}

		hex := currentLine[2][2 : len(currentLine[2])-1]
		length2 := parseIntHex(hex[:len(hex)-1])
		dir2Name := hex[len(hex)-1]
		dir2Coords := map[string][2]int{
			"0": {0, 1},
			"1": {1, 0},
			"2": {0, -1},
			"3": {-1, 0},
		}
		instructions = append(instructions, Instruction{
			direction:  directionNameToCoords[directionName],
			length:     parseInt(length),
			length2:    length2,
			direction2: dir2Coords[string(dir2Name)],
		})

	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error during input file read: %v\n", err.Error())
	}

	return instructions
}

func parseInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with: %s\n", s, err.Error())
	}

	return int(num)
}

func parseIntHex(s string) int {
	num, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		log.Fatalf("Failed to parse number %v with: %s\n", s, err.Error())
	}

	return int(num)
}

// func printMatrix(m [][]string) {
// 	fmt.Println("-----MATRIX START-----")
// 	for _, row := range m {
// 		fmt.Println(row)
// 	}
// 	fmt.Println("------MATRIX END------")
// }
