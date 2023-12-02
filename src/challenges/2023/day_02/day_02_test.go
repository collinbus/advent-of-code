package day_02

import (
	"advent-of-code/src/helpers"
	"fmt"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	sum := Day2Part1(input)

	fmt.Println(sum)
}
