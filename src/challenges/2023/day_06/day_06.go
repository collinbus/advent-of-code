package day_06

import (
	"regexp"
	"strconv"
)

func Day6Part1(input []string) int {
	times := readNumbers(input[0])
	distances := readNumbers(input[1])
	numberOfWays := 1
	for i, _ := range times {
		numberOfWays *= getNumberOfWays(times[i], distances[i])
	}
	return numberOfWays
}

func getNumberOfWays(time int, distance int) int {
	for currentTime := 2; currentTime <= time; currentTime++ {
		chargingSpeed := currentTime - 1
		max := chargingSpeed * (time - chargingSpeed)
		if max >= distance {
			return (time - 1) - ((chargingSpeed - 1) * 2)
		}
	}
	return 0
}

func readNumbers(input string) []int {
	numberRegex := regexp.MustCompile("[0-9]+")
	numbersAsString := numberRegex.FindAllString(input, -1)
	numbers := make([]int, len(numbersAsString))
	for i, s := range numbersAsString {
		number, _ := strconv.Atoi(s)
		numbers[i] = number
	}
	return numbers
}
