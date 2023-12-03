package day_03

import (
	"fmt"
	"strconv"
)

type Position struct {
	x int
	y int
}

func Day3Part1(input []string) int {
	sum := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '.' || isNumber(input[y][x]) {
				continue
			}
			indicesToExtract := getIndicesToExtract(input, x, y)
			numbers := extractNumbers(input, indicesToExtract, x, y)
			for _, number := range numbers {
				sum += number
			}
		}
	}
	return sum
}

func Day3Part2(input []string) int {
	sum := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] != '*' {
				continue
			}
			indicesToExtract := getIndicesToExtract(input, x, y)
			numbers := extractNumbers(input, indicesToExtract, x, y)
			if len(numbers) != 2 {
				continue
			}
			sum += numbers[0] * numbers[1]
		}
	}
	return sum
}

func extractNumbers(input []string, positions []*Position, x int, y int) []int {
	visited := map[string]int{}
	numbers := make([]int, 0)
	for _, position := range positions {
		if position.y == y {
			if position.x == x+1 {
				// read number to right
				number := readNumber(input, position.x, position.y, visited)
				numbers = append(numbers, number)
			}
			if position.x == x-1 {
				// look for number
				startIndex := lookForStartIndex(input, position.x, position.y)
				number := readNumber(input, startIndex, position.y, visited)
				numbers = append(numbers, number)
			}
		} else if visited[asString(position)] != 1 {
			startIndex := lookForStartIndex(input, position.x, position.y)
			number := readNumber(input, startIndex, position.y, visited)
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func lookForStartIndex(input []string, x int, y int) int {
	currentIndex := x
	for isNumber(input[y][currentIndex-1]) {
		currentIndex--
		if currentIndex-1 < 0 {
			return currentIndex
		}
	}
	return currentIndex
}

func readNumber(input []string, x int, y int, visited map[string]int) int {
	numberVal := ""
	current := input[y][x]
	for isNumber(current) {
		numberVal += string(input[y][x])
		visited[asString(&Position{
			x: x,
			y: y,
		})] = 1
		x++
		if x == len(input[y]) {
			break
		}
		current = input[y][x]
	}
	number, _ := strconv.Atoi(numberVal)
	return number
}

func getIndicesToExtract(input []string, x int, y int) []*Position {
	indicesToExtract := make([]*Position, 0)
	//check horizontal
	if isNumber(input[y][x-1]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x - 1,
			y: y,
		})
	}
	if isNumber(input[y][x+1]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x + 1,
			y: y,
		})
	}
	// check upper
	if isNumber(input[y-1][x]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x,
			y: y - 1,
		})
	}
	if isNumber(input[y-1][x-1]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x - 1,
			y: y - 1,
		})
	}
	if isNumber(input[y-1][x+1]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x + 1,
			y: y - 1,
		})
	}
	// check lower
	if isNumber(input[y+1][x]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x,
			y: y + 1,
		})
	}
	if isNumber(input[y+1][x-1]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x - 1,
			y: y + 1,
		})
	}
	if isNumber(input[y+1][x+1]) {
		indicesToExtract = append(indicesToExtract, &Position{
			x: x + 1,
			y: y + 1,
		})
	}
	return indicesToExtract
}

func isNumber(char uint8) bool {
	return char > 47 && char < 58
}

func asString(position *Position) string {
	return fmt.Sprintf("%d,%d", position.x, position.y)
}
