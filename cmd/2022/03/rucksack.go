package rucksack

import (
	"adventcode/utils"
	"fmt"
	"strconv"
)

// DecodeDay reads in an input file containing a list of random strings,
// where each string represents a 'rucksack':
//
// vJrwpWtwJgWrhcsFMMfFFhFp
//
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
//
// # PmmdzqPrVvPwwTWBwg
//
// Each rucksack has two compartments of equal length:
//
// vJrwpWtwJgWr        hcsFMMfFFhFp
//
// jqHRNqRjqzjGDLGL    rsFMfFZSrLrFZsSL
//
// # PmmdzqPrV           vPwwTWBwg
//
// Each letter in each compartment represents an 'item' and we need to find
// the first item that exist in both compartments:
//
// vJrw [p] WtwJgWr          hcsFMMfFFhF [p]           ---> 'p' exists in both compartments
//
// jqHRNqRjqzjGD [L] G [L]   rsFMfFZSr [L] rFZsS [L]   ---> 'L' exists in both compartments
//
// [P] mmdzq [P] rV          v [P] wwTWBwg             ---> 'P' exists in both compartments
//
// Each item has a unique value:
//
// 'a' through 'z' have values 1 through 26.
//
// 'A' through 'Z' have values 27 through 52.
//
// In part 1, we need to find the total values for all the common
// items that exist in both compartments.
//
// In part 2, every set of three rucksacks represents a group
// and we need to find the total values for all the common items
// that exist in each group.
func DecodeDay(path string) error {
	// Read input file.
	lines, err := utils.ReadLines(path)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Solve part 1: Get total value.
	solution, err := getCompartmentSolution(lines)
	if err != nil {
		return fmt.Errorf("failed to get solution for part 1: %v", err)
	}

	expected := expectedSolution(1)
	if solution != expected {
		return fmt.Errorf("unexpected solution found for part 1: %d != %d", solution, expected)
	}

	// Solve part 2: Get total score again but must choose weapon to get specified outcome.
	solution, err = getGroupSolution(lines)
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
	solution, err := utils.GetPuzzleSolution("2022/03/puzzle.txt", part)
	if err != nil {
		return -1
	}

	expected, err := strconv.Atoi(solution)
	if err != nil {
		return -1
	}

	return expected
}

// findCommonItem returns the value of a common item in two slices.
func findCommonItem(set1, set2 string) int {
	for _, item1 := range set1 {
		for _, item2 := range set2 {
			if item1 == item2 {
				return int(item1)
			}
		}
	}

	return 0
}

// findCommonItemGroup returns the value of a common item in three slices.
func findCommonItemGroup(set1, set2, set3 string) int {
	for _, item1 := range set1 {
		for _, item2 := range set2 {
			if item1 == item2 {
				for _, item3 := range set3 {
					if item1 == item3 {
						return int(item1)
					}
				}
			}
		}
	}

	return 0
}

// getItemValue converts the item's ascii value to the actual
// value we want to sum.
// 'a' through 'z' have ascii values 1 through 26.
// 'A' through 'Z' have values 27 through 52.
func convertItemValue(item int) int {
	if item <= 0 {
		return 0
	}

	if item >= 97 && item <= 122 {
		// 'a' through 'z'
		return item - 96
	}

	// 'A' through 'Z'
	return item - 38
}

// getCompartmentSolution returns the total values for all the common
// items that exist in both compartments for each rucksack.
func getCompartmentSolution(rucksacks []string) (int, error) {
	var total int
	for _, rucksack := range rucksacks {
		// Split rucksack into two compartments.
		compartment1 := rucksack[0 : len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]

		// Search for common item between each compartment.
		item := findCommonItem(compartment1, compartment2)

		// Update total amount.
		total += convertItemValue(item)
	}

	// Return total value.
	return total, nil
}

// getGroupSolution returns the total total values for all the common
// items that exist in each rucksack grouping.
func getGroupSolution(rucksacks []string) (int, error) {
	var total int
	for i := 0; i < len(rucksacks); i += 3 {
		// Get a group of three rucksacks.
		rucksack1 := rucksacks[i]
		rucksack2 := rucksacks[i+1]
		rucksack3 := rucksacks[i+2]

		// Search for common item between all three rucksacks.
		item := findCommonItemGroup(rucksack1, rucksack2, rucksack3)

		// Update total amount.
		total += convertItemValue(item)
	}

	// Return total value.
	return total, nil
}
