package main

import (
	calendar2022 "adventcode/cmd/2022"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"time"
)

func main() {
	// Track time to process all advent calenders.
	start := time.Now()

	// Get sub directories from current directory where each one represents
	// an advent calendar year.
	years, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("failed to read current directory: %v\n", err)
		return
	}

	// Decode each calendar year's puzzles.
	for _, year := range years {
		if !isYearDir(year) {
			continue
		}

		// Display year currently decoding.
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Printf("                                   %s                                       \n", year.Name())
		fmt.Println("--------------------------------------------------------------------------------")

		// Decode current year's calendar puzzles.
		switch year.Name() {
		case "2022":
			calendar2022.DecodeYear()
		}
	}

	// Calculate total execution time.
	duration := time.Since(start)

	// Display total time.
	fmt.Printf("TOTAL EXECUTION TIME: %12s\n", duration)
}

// isYearDir checks if the passed item is a directory and is named
// with a year. Must be a directory with a four digit name.
func isYearDir(dir fs.DirEntry) bool {
	if !dir.IsDir() {
		return false
	}

	name := dir.Name()
	if len(name) != 4 {
		return false
	}

	if _, err := strconv.ParseInt(name, 10, 64); err != nil {
		return false
	}

	// Is a year dir
	return true
}
