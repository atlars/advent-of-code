package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	day06Part2()
}

//go:embed input.txt
var input string

// regex to detect blocks of text divided by an arbitrary length of whitespaces
var whiteSpaceRegex = regexp.MustCompile(`\s+`)

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
	rows := strings.Split(input, "\n")
	operators := strings.Split(whiteSpaceRegex.ReplaceAllString(strings.TrimSpace(rows[len(rows)-1]), " "), " ")

	sum := 0

	line := 0
	for col := 0; col < len(operators); col++ {
		var nums []int
		hasNextNumber := true
		for hasNextNumber {
			num := 0
			for row := 0; row < len(rows)-1; row++ {
				if line < len(rows[row]) && rows[row][line] != ' ' {
					// append digit to the end of the number
					num *= 10
					num += int(rows[row][line] - '0')
				}
			}
			line++
			if num > 0 {
				nums = append(nums, num)
			} else {
				hasNextNumber = false
			}
		}
		fmt.Printf("%d numbers: %v op: %s\n", col, nums, operators[col])

		// calculate parsed numbers according to the given math operator
		result := nums[0]
		for i := 1; i < len(nums); i++ {
			switch operators[col] {
			case "+":
				result += nums[i]
			case "*":
				result *= nums[i]
			default:
				log.Fatalf("unknown operator %s", operators[col])
			}
		}
		sum += result
	}
	fmt.Printf("final sum: %d\n", sum)
}
