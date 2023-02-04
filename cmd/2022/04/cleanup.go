package cleanup

import (
	"adventcode/utils"
	"fmt"
	"strconv"
)

// DecodeDay reads in an input file containing a list of
func DecodeDay(path string) error {
	// Read input file.
	lines, err := utils.ReadLines(path)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Solve part 1: Get total score.
	solution, err := getSolution(lines)
	if err != nil {
		return fmt.Errorf("failed to get solution for part 1: %v", err)
	}

	expected := expectedSolution(1)
	if solution != expected {
		return fmt.Errorf("unexpected solution found for part 1: %d != %d", solution, expected)
	}

	// Solve part 2: Get total score again but must choose weapon to get specified outcome.
	solution, err = getSolution(lines)
	if err != nil {
		return fmt.Errorf("failed to get solution for part 2: %v", err)
	}

	expected = expectedSolution(2)
	if solution != expected {
		return fmt.Errorf("unexpected solution found for part 2: %d != %d", solution, expected)
	}

	return nil
}

// expectedSolution return the expected solution for specified part.
func expectedSolution(part int) int {
	solution, err := utils.GetPuzzleSolution("2022/04/puzzle.txt", part)
	if err != nil {
		return -1
	}

	expected, err := strconv.Atoi(solution)
	if err != nil {
		return -1
	}

	return expected
}

func getSolution(lines []string) (int, error) {
	var total int
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	// Return total points.
	return total, nil
}
