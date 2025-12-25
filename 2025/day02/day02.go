package main

import (
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	day02Part2()
}

func day02() {
	filePath, err := filepath.Abs("./2025/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	input := string(fileBytes)
	ranges := strings.Split(input, ",")

	var ids []string
	for _, stringRange := range ranges {
		r := strings.Split(stringRange, "-")

		intRange := generateRange(r[0], r[1])
		ids = append(ids, intRange...)
	}

	invalidSum := 0
	// Check for invalid ids
	for _, id := range ids {
		length := len(id)
		if length%2 == 0 && id[:length/2] == id[length/2:length] {
			number, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			invalidSum += number
		}
	}
	log.Printf("sum of invalid ids: %d", invalidSum)
}

func day02Part2() {
	filePath, err := filepath.Abs("./2025/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	input := string(fileBytes)
	ranges := strings.Split(input, ",")

	var ids []string
	for _, stringRange := range ranges {
		r := strings.Split(stringRange, "-")

		intRange := generateRange(r[0], r[1])
		ids = append(ids, intRange...)
	}

	invalidSum := 0
	// Check for invalid ids
	for _, id := range ids {
		if hasRepeatedSequence(id) {
			number, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			invalidSum += number
		}
	}
	log.Printf("sum of invalid ids: %d", invalidSum)
}

func hasRepeatedSequence(id string) bool {
	if len(id) <= 1 {
		return false
	}
	if len(id)%2 != 0 && strings.Repeat(string(id[0]), len(id)) == id {
		return true
	}

	for i := 1; i <= len(id)/2; i++ {
		subSequence := id[:i]
		repeat := len(id) / i
		if strings.Repeat(subSequence, repeat) == id {
			return true
		}
	}

	return false
}

// generateRange generates an incremental range of numbers returned as strings
func generateRange(start, end string) []string {
	startRange, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}
	endRange, err := strconv.Atoi(end)
	if err != nil {
		log.Fatal(err)
	}
	capacity := int(math.Abs(float64(startRange-endRange))) + 1
	result := make([]string, 0, capacity)
	for i := startRange; i <= endRange; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}
