package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	day05Part2()
}

type Interval struct {
	start, end int
}

//go:embed input.txt
var input string

func day05Part1() {
	ranges := strings.Split(input, "\n\n")
	freshInput := ranges[0]

	freshRanges := make([]Interval, len(freshInput))

	for index, rangeText := range strings.Split(freshInput, "\n") {
		r := strings.Split(rangeText, "-")
		start, err := strconv.Atoi(r[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatal(err)
		}
		freshRanges[index] = Interval{start, end}
	}

	availableInput := strings.Split(ranges[1], "\n")
	availableIds := make([]int, len(availableInput))
	for index, idText := range availableInput {
		id, err := strconv.Atoi(idText)
		if err != nil {
			log.Fatal(err)
		}
		availableIds[index] = id
	}

	countFreshIngredients := 0
	for _, availableId := range availableIds {
		if slices.ContainsFunc(freshRanges, inRange(availableId)) {
			countFreshIngredients++
		}

	}

	fmt.Printf("number of fresh ingredients that are available: %d\n", countFreshIngredients)
}

// inRange checks if a given id is in the range of an Interval
func inRange(id int) func(Interval) bool {
	return func(interval Interval) bool {
		if id >= interval.start && id <= interval.end {
			return true
		}
		return false
	}
}

// idea: merge ranges, so that no overlapping range exist anymore
func day05Part2() {
	freshInput := strings.Split(input, "\n\n")[0]
	rangesInput := strings.Split(freshInput, "\n")

	ranges := make([]Interval, 0, len(rangesInput))

	// extract ranges from input text
	for _, rangeText := range rangesInput {
		r := strings.Split(rangeText, "-")
		start, err := strconv.Atoi(r[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, Interval{start, end})
	}

	// sort ranges by start
	slices.SortFunc(ranges, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})

	mergedRanges := make([]Interval, 0)

	// merge ranges to remove overlapping ranges
	curRange := ranges[0]
	for i := 1; i < len(ranges); i++ {
		nextRange := ranges[i]
		// ranges overlap
		if nextRange.start <= curRange.end {
			curRange.end = max(curRange.end, nextRange.end)
		} else {
			mergedRanges = append(mergedRanges, curRange)
			curRange = nextRange
		}
	}
	mergedRanges = append(mergedRanges, curRange)

	count := 0
	for _, mergedRange := range mergedRanges {
		count += mergedRange.end - mergedRange.start + 1 // +1 because intervals are inclusive
	}

	fmt.Printf("merged ranges: %v\n", mergedRanges)
	fmt.Printf("count of merged ranges: %v\n", count)
}
