package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	day06Part2()
}

//go:embed input.txt
var input string

func day06Part1() {
	matrix, operators := parseMatrix(input)

	// calculate terms
	sum := 0
	for x := 0; x < len(matrix[0]); x++ {
		result := matrix[0][x]
		for y := 1; y < len(matrix); y++ {
			num := matrix[y][x]
			switch operators[x] {
			case "+":
				result += num
			case "*":
				result *= num
			default:
				log.Fatalf("unknown operator %s", operators[x])
			}
		}
		sum += result
	}
	fmt.Printf("sum: %d\n", sum)
}

// parseMatrix parses the numbers and operators of the input string
func parseMatrix(input string) (matrix [][]int, operators []string) {
	rows := strings.Split(input, "\n")
	matrix = make([][]int, len(rows)-1)
	whiteSpaceRegex := regexp.MustCompile(`\s+`)
	for i := 0; i < len(rows)-1; i++ {
		// replace multiple whitespaces to a single whitespace
		trimmedRow := whiteSpaceRegex.ReplaceAllString(strings.TrimSpace(rows[i]), " ")
		numbers := strings.Split(trimmedRow, " ")

		// parse numbers into row
		matrix[i] = make([]int, len(numbers))
		for index, textNumber := range numbers {
			number, err := strconv.Atoi(textNumber)
			if err != nil {
				log.Fatal(err)
			}
			matrix[i][index] = number
		}
	}
	// parse math operators
	trimmedOps := whiteSpaceRegex.ReplaceAllString(strings.TrimSpace(rows[len(rows)-1]), " ")
	operators = strings.Split(trimmedOps, " ")
	return
}

func day06Part2() {

}

// digits extracts all digits of a number
func digits(n int) []int {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return []int{0}
	}
	var res []int
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	slices.Reverse(res)
	return res
}
