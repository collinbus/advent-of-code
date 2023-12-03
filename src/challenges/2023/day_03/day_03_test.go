package day_03

import (
	"advent-of-code/src/helpers"
	"fmt"
	"testing"
)

func TestDay3Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day3Part1(input)

	fmt.Println(sum)
}

func TestDay3Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day3Part2(input)

	fmt.Println(sum)
}
