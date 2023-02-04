package calendar2022

import (
	calories "adventcode/cmd/2022/01"
	game "adventcode/cmd/2022/02"
	rucksack "adventcode/cmd/2022/03"
	cleanup "adventcode/cmd/2022/04"
	"adventcode/utils"
	"fmt"
	"time"
)

func DecodeYear() {
	for day := 1; day < 26; day++ {
		// Track time.
		start := time.Now()

		// Get puzzle title.
		title, err := utils.GetPuzzleTitle(fmt.Sprintf("2022/%02d/puzzle.txt", day))
		if err != nil {
			fmt.Printf("failed to get puzzle title for day %02d: %v", day, err)
			continue
		}

		// Get puzzle input file path.
		path := fmt.Sprintf("2022/%02d/input.txt", day)

		// Decode day's puzzle.
		switch day {
		case 1:
			err = calories.DecodeDay(path)
		case 2:
			err = game.DecodeDay(path)
		case 3:
			err = rucksack.DecodeDay(path)
		case 4:
			err = cleanup.DecodeDay(path)
		default:
			err = fmt.Errorf("not implemented")
		}

		// Calculate execution time.
		duration := time.Since(start)

		// Display results.
		if err != nil {
			fmt.Printf("DAY %02d: [%12s]\t[TODO][%s][ERROR: %v]\n", day, duration, title, err)
		} else {
			fmt.Printf("DAY %02d: [%12s]\t[SOLVED][%s]\n", day, duration, title)
		}
	}
}
