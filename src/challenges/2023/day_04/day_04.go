package day_04

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day4Part1(input []string) int {
	score := 0
	for _, line := range input {
		splittedInput := strings.Split(line, ": ")
		winningNumbers := getScoresFrom(splittedInput[1], 0)
		handNumbers := getScoresFrom(splittedInput[1], 1)

		sort.Ints(winningNumbers)
		sort.Ints(handNumbers)

		roundScore := 0
		winningIndex := 0
		handIndex := 0
		for handIndex < len(handNumbers) {
			winningNumber := winningNumbers[winningIndex]
			handNumber := handNumbers[handIndex]
			if handNumber == winningNumber {
				roundScore = increaseScore(roundScore)
			}
			if handNumber >= winningNumber {
				winningIndex++
			}
			if handNumber < winningNumber {
				handIndex++
			}
			if winningIndex >= len(winningNumbers) {
				break
			}
		}
		score += roundScore
	}
	return score
}

func Day4Part2(input []string) int {
	copies := makeCopiesMap(input)
	for i, line := range input {
		numberOfCopies := copies[i]

		for j := 0; j < numberOfCopies; j++ {
			splittedInput := strings.Split(line, ": ")
			winningNumbers := getScoresFrom(splittedInput[1], 0)
			handNumbers := getScoresFrom(splittedInput[1], 1)

			sort.Ints(winningNumbers)
			sort.Ints(handNumbers)

			winningIndex := 0
			handIndex := 0
			nextCopyIndex := i + 1
			for handIndex < len(handNumbers) {
				winningNumber := winningNumbers[winningIndex]
				handNumber := handNumbers[handIndex]
				if handNumber == winningNumber {
					if nextCopyIndex < len(input) {
						copies[nextCopyIndex] += 1
						nextCopyIndex++
					}
				}
				if handNumber >= winningNumber {
					winningIndex++
				}
				if handNumber < winningNumber {
					handIndex++
				}
				if winningIndex >= len(winningNumbers) {
					break
				}
			}
		}
	}
	return copyCount(copies)
}

func copyCount(copies map[int]int) int {
	totalCopies := 0
	for _, numberOfCopies := range copies {
		totalCopies += numberOfCopies
	}
	return totalCopies
}

func makeCopiesMap(input []string) map[int]int {
	copies := make(map[int]int, len(input))
	for i := 0; i < len(input); i++ {
		copies[i] = 1
	}
	return copies
}

func increaseScore(score int) int {
	if score == 0 {
		return 1
	}
	return score * 2
}

func getScoresFrom(input string, numbersIndex int) []int {
	numberRegex := regexp.MustCompile("[0-9]+")
	numbersAsString := strings.Split(input, " | ")
	numbers := numberRegex.FindAllString(numbersAsString[numbersIndex], -1)
	return toInts(numbers)
}

func toInts(stringValues []string) []int {
	ints := make([]int, len(stringValues))
	for i, value := range stringValues {
		number, _ := strconv.Atoi(value)
		ints[i] = number
	}
	return ints
}
