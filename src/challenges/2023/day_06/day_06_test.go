package day_06

import (
	"advent-of-code/src/helpers"
	"fmt"
	"testing"
)

func TestDay6Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day6Part1(input)

	fmt.Println(sum)
}
