package calories

import (
	"adventcode/utils"
	"fmt"
	"sort"
	"strconv"
)

// DecodeDay reads in an input file containing a list of number groupings where
// each group represents the amount of calories per elf.
//
// For example:
//
// 1000
// 2000
// 3000 -> elf 1 consumed 6000 calories
//
// 4000 -> elf 2 consumed 4000 calories
//
// 5000
// 6000 -> elf 3 consumed 11000 calories
//
// 7000
// 8000
// 9000 -> elf 4 consumed 24000 calories
//
// 10000 -> elf 5 consumed 10000 calories
//
// In part 1, we find which elf is carrying the most calories and what the total
// calories are.
//
// In part 2, we find the top three elves carrying the most calories and what
// the total of all their calories are.
func DecodeDay(path string) error {
	// Read input file.
	lines, err := utils.ReadLines(path)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Get list of elves' total calories.
	elves, err := getElves(lines)
	if err != nil {
		return fmt.Errorf("failed to get elves: %v", err)
	}

	// Solve part 1: Top elf.
	solution := elves[0]
	expected := expectedSolution(1)
	if solution != expected {
		return fmt.Errorf("unexpected solution found for part 1: %d != %d", solution, expected)
	}

	// Solve part 2: Top three elves.
	solution = elves[0] + elves[1] + elves[2]
	expected = expectedSolution(2)
	if solution != expected {
		return fmt.Errorf("unexpected solution found for part 2: %d != %d", solution, expected)
	}

	return nil
}

// expectedSolution return the expected solution for specified part.
func expectedSolution(part int) int {
	solution, err := utils.GetPuzzleSolution("2022/01/puzzle.txt", part)
	if err != nil {
		return -1
	}

	expected, err := strconv.Atoi(solution)
	if err != nil {
		return -1
	}

	return expected
}

// getElves returns an ordered list of elves and their total calories.
// The elves with the most calories will be at the front of the list.
func getElves(lines []string) ([]int, error) {
	var elves []int

	var total int
	for _, line := range lines {
		switch len(line) {
		case 0:
			// We have finished an elve's group of calories so add
			// total calories to list.
			elves = append(elves, total)

			// Reset for next elf grouping.
			total = 0
		default:
			// Convert line to calories.
			calories, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("failed to get calories from input line %q: %v", line, err)
			}

			// Update total calories.
			total += calories
		}
	}

	// Add last total calories to list.
	elves = append(elves, total)

	// Sort list so that the elves with the highest calories will
	// be in the front.
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	return elves, nil
}
