package day_01

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getNumberFrom(line, false)
	}
	return sum
}

func Day1Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getNumberFrom(line, true)
	}
	return sum
}

func getNumberFrom(line string, readFullDigits bool) int {
	first := -1
	last := -1
	start := 0
	end := len(line) - 1
	for first == -1 || last == -1 {
		if first == -1 {
			first = readDigit(line, start, readFullDigits)
		}
		if last == -1 {
			last = readDigit(line, end, readFullDigits)
		}
		start++
		end--
	}
	numberAsString := fmt.Sprintf("%d%d", first, last)
	number, _ := strconv.Atoi(numberAsString)
	return number
}

func readDigit(line string, index int, readFullDigits bool) int {
	char := line[index]
	digit := getNumberValue(char)
	if digit != -1 || !readFullDigits {
		return digit
	}
	inputString := line[index:]
	translation := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for key, _ := range translation {
		if strings.Index(inputString, key) == 0 {
			return translation[key]
		}
	}
	return -1
}

func getNumberValue(char uint8) int {
	if char > 47 && char < 58 {
		return int(char - 48)
	}
	return -1
}
