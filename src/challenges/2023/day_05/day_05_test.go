package day_05

import (
	"advent-of-code/src/helpers"
	"fmt"
	"testing"
)

func TestDay5Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day5Part1(input)

	fmt.Println(sum)
}

func TestDay5Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day5Part2(input)

	fmt.Println(sum)
}
