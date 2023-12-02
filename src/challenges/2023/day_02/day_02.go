package day_02

import (
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	rounds []*Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func Day2Part1(input []string) int {
	reds := 12
	greens := 13
	blues := 14
	sumIds := 0
	for _, line := range input {
		game := parseGame(line)
		lost := false
		for _, round := range game.rounds {
			reds -= round.red
			greens -= round.green
			blues -= round.blue
			if reds < 0 || greens < 0 || blues < 0 {
				lost = true
			}
			reds += round.red
			greens += round.green
			blues += round.blue
			if lost {
				break
			}
		}
		if !lost {
			sumIds += game.id
		}
	}
	return sumIds
}

func parseGame(input string) *Game {
	splittedGame := strings.Split(input, ": ")
	id := extractNumber(splittedGame[0], "[0-9]+")
	roundsAsString := strings.Split(splittedGame[1], "; ")
	game := &Game{
		id:     id,
		rounds: make([]*Round, len(roundsAsString)),
	}
	for i, roundAsString := range roundsAsString {
		movesAsString := strings.Split(roundAsString, ", ")
		round := &Round{}
		for _, moveAsString := range movesAsString {
			number := extractNumber(moveAsString, "[0-9]+")
			if strings.Contains(moveAsString, "green") {
				round.green = number
			}
			if strings.Contains(moveAsString, "blue") {
				round.blue = number
			}
			if strings.Contains(moveAsString, "red") {
				round.red = number
			}
		}
		game.rounds[i] = round
	}
	return game
}

func extractNumber(input string, pattern string) int {
	numberRegex := regexp.MustCompile(pattern)
	idAsString := numberRegex.FindString(input)
	number, _ := strconv.Atoi(idAsString)
	return number
}
