package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"path/filepath"
)

func main() {
	day03Part2()
}

func day03() {
	path, err := filepath.Abs("./2025/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalOutputJoltage := 0
	for scanner.Scan() {
		bank := scanner.Text()

		firstIndex := 0
		for i := 0; i < len(bank)-1; i++ {
			if int(bank[i]) > int(bank[firstIndex]) {
				firstIndex = i
			}
		}

		secondIndex := firstIndex + 1
		for i := secondIndex; i < len(bank); i++ {
			if int(bank[i]) > int(bank[secondIndex]) {
				secondIndex = i
			}
		}

		bankOutputJoltage := (bank[firstIndex]-'0')*10 + bank[secondIndex] - '0'
		totalOutputJoltage += int(bankOutputJoltage)

		log.Printf("bank output joltage %d", bankOutputJoltage)
	}
	log.Printf("total output joltage %d", totalOutputJoltage)
}

func day03Part2() {
	path, err := filepath.Abs("./2025/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxJoltage := 0
	for scanner.Scan() {
		bank := scanner.Text()
		indexes := make([]int, 12)

		for i := 0; i < len(indexes); i++ {
			start := 0
			if i > 0 {
				start = indexes[i-1] + 1
				indexes[i] = start
			}

			// find the highest number between the last number and the enough digits left
			for j := start; j <= len(bank)-12+i; j++ {
				if int(bank[j]) > int(bank[indexes[i]]) {
					indexes[i] = j
				}
			}
		}

		joltage := 0
		for i := len(indexes) - 1; i >= 0; i-- {
			joltage += int(bank[indexes[i]]-'0') * int(math.Pow10(len(indexes)-i-1))
		}

		maxJoltage += joltage

	}
	log.Printf("max joltage: %d", maxJoltage)
}
