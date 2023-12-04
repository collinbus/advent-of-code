package day_04

import (
	"advent-of-code/src/helpers"
	"fmt"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day4Part1(input)

	fmt.Println(sum)
}

func TestDay4Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day4Part2(input)

	fmt.Println(sum)
}
