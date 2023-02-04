package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
