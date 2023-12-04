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
		numbersAsString := strings.Split(splittedInput[1], " | ")
		numberRegex := regexp.MustCompile("[0-9]+")
		winningNumbersAsString := numberRegex.FindAllString(numbersAsString[0], -1)
		handNumbersAsString := numberRegex.FindAllString(numbersAsString[1], -1)
		winningNumbers := toInts(winningNumbersAsString)
		handNumbers := toInts(handNumbersAsString)
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

func increaseScore(score int) int {
	if score == 0 {
		return 1
	}
	return score * 2
}

func toInts(stringValues []string) []int {
	ints := make([]int, len(stringValues))
	for i, value := range stringValues {
		number, _ := strconv.Atoi(value)
		ints[i] = number
	}
	return ints
}
