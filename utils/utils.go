package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadLines accepts a path to a file, opens that file for reading
// and returns a slice of strings from the contents of the file.
func ReadLines(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// GetPuzzleTitle accepts a path to a puzzle description file, opens
// file for reading and searches for the title of the puzzle
// file describes. For example, it will search for the string:
//
// # --- Day 1: Calorie Counting ---
//
// will return:
//
// # Calorie Counting
func GetPuzzleTitle(path string) (string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %v", err)
	}

	for _, line := range lines {
		if strings.Contains(line, "--- Day") {
			title := strings.Split(line, ": ")[1]
			title = strings.Split(title, " ---")[0]
			return title, nil
		}
	}

	return "", fmt.Errorf("title not found in input file %q: %v", path, err)
}

// GetPuzzleTitle accepts a path to a puzzle description file, opens
// file for reading and searches for the solution of the puzzle
// file describes if the solution has already been found.
// For example, it will search for the string:
//
// # Your puzzle answer was 12345
//
// and return:
//
// # 12345
func GetPuzzleSolution(path string, part int) (string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %v", err)
	}

	count := 0
	for _, line := range lines {
		if strings.Contains(line, "Your puzzle answer was ") {
			count++
			if count == part {
				solution := strings.Split(line, "Your puzzle answer was ")[1]
				solution = strings.Split(solution, ".")[0]
				return solution, nil
			}
		}
	}

	return "", fmt.Errorf("solution not found in input file for part %d: %v", part, err)
}
