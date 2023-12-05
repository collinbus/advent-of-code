package day_05

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type GardenMap struct {
	from   string
	to     string
	ranges []*Range
}

type Range struct {
	destinationStart int
	sourceStart      int
	length           int
}

func NewRange(destinationStart int, sourceStart int, length int) *Range {
	return &Range{destinationStart: destinationStart, sourceStart: sourceStart, length: length}
}

func Day5Part1(input []string) int {
	seeds := readSeeds(input[0])
	maps := readMaps(input[2:])
	for _, currentMap := range maps {
		sort.Ints(seeds)
		sort.Slice(currentMap.ranges, func(i, j int) bool {
			iRange := currentMap.ranges[i]
			jRange := currentMap.ranges[j]
			return iRange.sourceStart < jRange.sourceStart
		})
		seedIndex := 0
		rangeIndex := 0
		for seedIndex < len(seeds) && rangeIndex < len(currentMap.ranges) {
			currentSeed := seeds[seedIndex]
			currentRange := currentMap.ranges[rangeIndex]
			if currentSeed >= currentRange.sourceStart && currentSeed < currentRange.sourceStart+currentRange.length {
				seeds[seedIndex] = currentRange.destinationStart + (currentSeed - currentRange.sourceStart)
				seedIndex++
			} else {
				rangeIndex++
			}
		}
	}
	sort.Ints(seeds)
	return seeds[0]
}

func readMaps(input []string) []*GardenMap {
	maps := make([]*GardenMap, 0)
	var ranges []*Range
	var currentMap *GardenMap
	for _, line := range input {
		if strings.Contains(line, "-to-") {
			mapRegexp := regexp.MustCompile(`(\w+)-to-(\w+)`)
			fromTo := mapRegexp.FindStringSubmatch(line)
			currentMap = &GardenMap{
				from:   fromTo[1],
				to:     fromTo[2],
				ranges: nil,
			}
			ranges = make([]*Range, 0)
		} else if line == "" {
			currentMap.ranges = ranges
			maps = append(maps, currentMap)
		} else {
			values := strings.Split(line, " ")
			destination, _ := strconv.Atoi(values[0])
			source, _ := strconv.Atoi(values[1])
			length, _ := strconv.Atoi(values[2])
			currentRange := NewRange(destination, source, length)
			ranges = append(ranges, currentRange)
		}
	}
	currentMap.ranges = ranges
	maps = append(maps, currentMap)
	return maps
}

func readSeeds(input string) []int {
	seedsAsString := strings.Split(input[7:], " ")
	seeds := make([]int, len(seedsAsString))
	for i, value := range seedsAsString {
		seed, _ := strconv.Atoi(value)
		seeds[i] = seed
	}
	return seeds
}
