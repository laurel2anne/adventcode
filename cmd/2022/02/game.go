package game

import (
	"adventcode/utils"
	"fmt"
	"strconv"
)

// DecodeDay reads in an input file containing a list of two columns where
// the first column contains the letter A, B, or C. The second column contains
// letter A, B, or C:
//
// For example:
//
//	A Y
//	B X
//	C Z
//
// The first column represents the opponents weapon:
// A = Rock
// B = Paper
// C = Scissors
//
// Each weapon has a base score:
// Rock = 1
// Paper = 2
// Scissors = 3
//
// Each match then increments your score:
// Lost = 0
// Tie = 3
// Win = 6
//
// In part 1, we assume the second column represents the weapon we use
// and calculate my total score.
// X = Rock
// Y = Paper
// Z = Scissors
//
// In part 2, we find out that the second column actually represents
// the outcome of what we want to happen and calculate my total score.
//
// X = Lose
// Y = Tie
// Z = Win
func DecodeDay(path string) error {
	// Read input file.
	lines, err := utils.ReadLines(path)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Solve part 1: Get total score.
	solution, err := getSolution(lines, false)
	if err != nil {
		return fmt.Errorf("failed to get solution for part 1: %v", err)
	}

	expected := expectedSolution(1)
	if solution != expected {
		return fmt.Errorf("unexpected solution found for part 1: %d != %d", solution, expected)
	}

	// Solve part 2: Get total score again but must choose weapon to get specified outcome.
	solution, err = getSolution(lines, true)
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
	solution, err := utils.GetPuzzleSolution("2022/02/puzzle.txt", part)
	if err != nil {
		return -1
	}

	expected, err := strconv.Atoi(solution)
	if err != nil {
		return -1
	}

	return expected
}

func getSolution(lines []string, choose bool) (int, error) {
	var total int
	for _, line := range lines {
		// Get the opponents weapon.
		opponentWeapon := getWeapon(line[0:1])

		// Get my weapon depending on which part we are trying to solve.
		var myWeapon string
		if choose {
			myWeapon = chooseWeapon(line[2:3], opponentWeapon)
		} else {
			myWeapon = getWeapon(line[2:3])
		}

		// Get my score after match.
		score := getMatchScore(myWeapon, opponentWeapon)

		// Incrememnt total.
		total += score
	}

	// Return total points.
	return total, nil
}

// getWeapon return weapon type based on symbol.
func getWeapon(symbol string) string {
	switch symbol {
	case "A":
		return "Rock"
	case "B":
		return "Paper"
	case "C":
		return "Scissors"
	case "X":
		return "Rock"
	case "Y":
		return "Paper"
	case "Z":
		return "Scissors"
	}

	return ""
}

// chooseWeapon returns weapon based on opponents weapon
// and the outcome we want.
func chooseWeapon(outcome, weapon string) string {
	switch outcome {
	case "X":
		// Want to lose.
		switch weapon {
		case "Rock":
			return "Scissors"
		case "Paper":
			return "Rock"
		case "Scissors":
			return "Paper"
		}
	case "Y":
		// Want to tie.
		return weapon
	case "Z":
		// Want to win.
		switch weapon {
		case "Rock":
			return "Paper"
		case "Paper":
			return "Scissors"
		case "Scissors":
			return "Rock"
		}
	}

	return ""
}

// getWeaponScore returns the base score for a weapon.
func getWeaponScore(weapon string) int {
	switch weapon {
	case "Rock":
		return 1
	case "Paper":
		return 2
	case "Scissors":
		return 3
	}

	return 0
}

// getMatchScore returns the score for my weapon when matched
// with the opponent's weapon.
func getMatchScore(myWeapon, opponentWeapon string) int {
	// Get base score for my weapon
	myScore := getWeaponScore(myWeapon)

	if myWeapon == opponentWeapon {
		// In a tie, both get an extra 3 points
		myScore += 3
		return myScore
	}

	switch myWeapon {
	case "Rock":
		if opponentWeapon == "Scissors" {
			// Rock vs Scissors: I win an extra 6 points
			myScore += 6
		}
	case "Paper":
		if opponentWeapon == "Rock" {
			// Paper vs Rock: I win an extra 6 points
			myScore += 6
		}
	case "Scissors":
		if opponentWeapon == "Paper" {
			// Scissors vs Paper: I win an extra 6 points
			myScore += 6
		}
	}

	return myScore
}
