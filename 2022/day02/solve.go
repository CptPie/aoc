package day02

import (
	"errors"
	"fmt"
	"strings"
)

// Col 1:
// A - Rock
// B - Paper
// C - Scissors

// Col 2:
// X - Rock     - 1
// Y - Paper    - 2
// Z - Scissors - 3

// Win: 6, Draw: 3, Loss: 0

func Solve(fileContents []string) ([]int, error) {
	var results []int

	solution, err := solvePart1(fileContents)

	if err != nil {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

func solvePart1(fileContents []string) (int, error) {
	sum := 0
	for _, line := range fileContents {
		res, err := parsePart1(line)
		if err != nil {
			return 0, err
		}
		sum += res
	}
	return sum, nil
}

func solvePart2(fileContents []string) (int, error) {
	sum := 0
	for _, line := range fileContents {

		res, err := parsePart2(line)
		if err != nil {
			return 0, err
		}
		sum += res
	}
	return sum, nil
}

// Required        Opponent       Score
// X - Win   - 6    A - Rock     - 1
// Y - Draw  - 3    B - Paper    - 2
// Z - Loose - 0    C - Scissors - 3

func parsePart2(line string) (int, error) {
	score := 0
	parts := strings.Fields(line)

	switch parts[1] {
	case "X":
		// we loose
		score += 0
		if parts[0] == "A" {
			// Opponent plays Rock
			// We play Scissors
			score += 3
		} else if parts[0] == "B" {
			// Opponent plays Paper
			// We play Rock
			score += 1
		} else if parts[0] == "C" {
			// Opponent plays Scissors
			// We play Paper
			score += 2
		} else {
			return 0, errors.New("Invalid input")
		}
	case "Y":
		// Draw
		score += 3
		if parts[0] == "A" {
			// Opponent plays Rock
			// We play Rock
			score += 1
		} else if parts[0] == "B" {
			// Opponent plays Paper
			// We play Paper
			score += 2
		} else if parts[0] == "C" {
			// Opponent plays Scissors
			// We play Scissors
			score += 3
		} else {
			return 0, errors.New("Invalid input")
		}
	case "Z":
		// we win
		score += 6
		if parts[0] == "A" {
			// Opponent plays Rock
			// We play Paper
			score += 2
		} else if parts[0] == "B" {
			// Opponent plays Paper
			// We play Scissors
			score += 3
		} else if parts[0] == "C" {
			// Opponent plays Scissors
			// We play Rock
			score += 1
		} else {
			return 0, errors.New("Invalid input")
		}
	default:
		return 0, errors.New("Invalid input")
	}
	return score, nil
}

func parsePart1(line string) (int, error) {
	score := 0
	parts := strings.Fields(line)

	switch parts[1] {
	case "X":
		score += 1
		if parts[0] == "A" {
			score += 3
		} else if parts[0] == "B" {
			score += 0
		} else if parts[0] == "C" {
			score += 6
		} else {
			return 0, errors.New("Invalid input")
		}
	case "Y":
		score += 2
		if parts[0] == "A" {
			score += 6
		} else if parts[0] == "B" {
			score += 3
		} else if parts[0] == "C" {
			score += 0
		} else {
			return 0, errors.New("Invalid input")
		}
	case "Z":
		score += 3
		if parts[0] == "A" {
			score += 0
		} else if parts[0] == "B" {
			score += 6
		} else if parts[0] == "C" {
			score += 3
		} else {
			return 0, errors.New("Invalid input")
		}
	default:
		return 0, errors.New("Invalid input")
	}
	return score, nil
}
