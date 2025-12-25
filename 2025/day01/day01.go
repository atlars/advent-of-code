package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	day02Part2()
}

func day01() {
	path, err := filepath.Abs("./2025/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	dial := 50
	password := 0
	for scanner.Scan() {
		s := scanner.Text()
		direction := string(s[0])
		rotation, err := strconv.Atoi(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		if direction == "L" {
			dial = (dial - rotation) % 100
			if dial < 0 {
				dial = dial + 100
			}
		} else if direction == "R" {
			dial = (dial + rotation) % 100
		}

		if dial == 0 {
			password += 1
		}
	}
	log.Printf("Calculated password: %d", password)
}

func day02Part2() {
	path, err := filepath.Abs("./2025/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	dial := 50
	password := 0
	for scanner.Scan() {
		s := scanner.Text()
		direction := string(s[0])
		rotation, err := strconv.Atoi(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		password += rotation / 100
		remainingRotation := rotation % 100
		if direction == "L" {
			if dial != 0 && dial-remainingRotation <= 0 {
				password++
			}

			dial = (dial - rotation) % 100
			if dial < 0 {
				dial = dial + 100
			}
		} else if direction == "R" {
			if dial != 0 && dial+remainingRotation > 99 {
				password++
			}
			dial = (dial + rotation) % 100
		}

	}
	log.Printf("Calculated password: %d", password)
}
