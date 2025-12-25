package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const maxNeighbourRolesCount int = 3

type Point struct {
	x, y int
}

var offsets = []Point{
	{-1, 0},  // Left
	{-1, -1}, // Top-Left
	{0, -1},  // Up
	{1, -1},  // Top-Right
	{1, 0},   // Right
	{1, 1},   // Bottom-Right
	{0, 1},   // Down
	{-1, 1},  // Bottom-Left
}

func main() {
	day04Part2()
}

func day04Part1() {
	path, err := filepath.Abs("./2025/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	paperMap := strings.Split(string(input), "\n")
	inBounds := func(x, y int) bool {
		if x < 0 || y < 0 || y >= len(paperMap) || x >= len(paperMap[0]) {
			return false
		}
		return true
	}

	accessibleRoles := 0

	for x := 0; x < len(paperMap[0]); x++ {
		for y := 0; y < len(paperMap); y++ {
			if paperMap[y][x] != '@' {
				continue
			}
			rolesCounter := 0

			for _, offset := range offsets {
				nX := x + offset.x
				nY := y + offset.y
				if inBounds(nX, nY) && paperMap[nY][nX] == '@' {
					rolesCounter++
				}
			}
			if rolesCounter <= maxNeighbourRolesCount {
				accessibleRoles++
			}
		}
	}

	log.Printf("accessible roles: %d", accessibleRoles)
}

func day04Part2() {
	path, err := filepath.Abs("./2025/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	paperMap := strings.Split(string(input), "\n")
	inBounds := func(x, y int) bool {
		if x < 0 || y < 0 || y >= len(paperMap) || x >= len(paperMap[0]) {
			return false
		}
		return true
	}

	totalRolesRemoved := 0
	for {
		var rolesToRemove []Point

		for x := 0; x < len(paperMap[0]); x++ {
			for y := 0; y < len(paperMap); y++ {
				if paperMap[y][x] != '@' {
					continue
				}
				rolesCounter := 0

				for _, offset := range offsets {
					nX := x + offset.x
					nY := y + offset.y
					if inBounds(nX, nY) && paperMap[nY][nX] == '@' {
						rolesCounter++
					}
				}
				if rolesCounter <= maxNeighbourRolesCount {
					rolesToRemove = append(rolesToRemove, Point{x, y})
				}
			}
		}

		for _, p := range rolesToRemove {
			line := paperMap[p.y]
			paperMap[p.y] = line[:p.x] + "x" + line[p.x+1:]
		}

		totalRolesRemoved += len(rolesToRemove)

		if len(rolesToRemove) <= 0 {
			break
		}
	}

	log.Printf("total roles removed: %d", totalRolesRemoved)
}
