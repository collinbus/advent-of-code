package day_01

import (
	"advent-of-code/src/helpers"
	"fmt"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day1Part1(input)

	fmt.Println(sum)
}

func TestDay1Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day1Part2(input)

	fmt.Println(sum)
}
