package day03

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Solve(fileContents []string) ([]string, error) {
	var results []string

	solution, err := solvePart1(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

func solvePart1(fileContents []string) (int, error) {
	sum := 0
	for _, line := range fileContents {
		value, err := day3CalculateLinePart1(line)

		if err != nil {
			return 0, err
		}

		sum += value
	}

	return sum, nil
}
func day3CalculateLinePart1(line string) (int, error) {

	if len(line)%2 != 0 {
		return 0, errors.New("Line is not even")
	}

	compartment1 := line[0 : len(line)/2]
	compartment2 := line[len(line)/2:]

	for _, character := range compartment1 {
		if strings.Contains(compartment2, string(character)) {
			return day3CalculateRuneScore(character), nil
		}
	}

	return 0, nil
}

func solvePart2(fileContents []string) (int, error) {
	sum := 0
	lines := []string{}
	if len(fileContents)%3 != 0 {
		return 0, errors.New("Input file not divisible by 3")
	}
	for _, line := range fileContents {
		lines = append(lines, line)
		if len(lines) == 3 {
			// third pack of the group
			value, err := day3CalculateValuePart2(lines)

			if err != nil {
				return 0, err
			}

			sum += value
			lines = []string{}
		}
	}

	return sum, nil
}

func day3CalculateValuePart2(lines []string) (int, error) {
	if len(lines) != 3 {
		return 0, fmt.Errorf("We've done fucked up")
	}
	for _, character := range lines[0] {
		if strings.Contains(lines[1], string(character)) && strings.Contains(lines[2], string(character)) {
			return day3CalculateRuneScore(character), nil
		}

	}
	return 0, fmt.Errorf("No match found")
}

func day3CalculateRuneScore(r rune) int {
	if unicode.IsLower(r) {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
}
