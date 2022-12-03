package day03

import (
	"2022/utils"
	"fmt"
	"strings"
	"unicode"
)

func Solve() ([]int, error) {
	day := 3

	results := []int{}
	
	solution, err := solvePart1(day)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(day)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

func solvePart1(day int) (int, error) {
	fileScanner, err := utils.ReadFile(day)
	_, _ = fileScanner, err

	if err != nil {
		return 0, err
	}

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		value, err := day3CalculateLinePart1(line)

		if err != nil {
			return 0, err
		}
		sum += value
	}

	return sum, nil
}
func day3CalculateLinePart1(line string) (int, error) {

	compartment1 := line[0 : len(line)/2]
	compartment2 := line[len(line)/2:]

	for _, character := range compartment1 {
		if strings.Contains(compartment2, string(character)) {
			return day3CalculateRuneScore(character), nil
		}
	}

	return 0, nil
}

func solvePart2(day int) (int, error) {

	fileScanner, err := utils.ReadFile(day)
	_, _ = fileScanner, err

	if err != nil {
		return 0, err
	}

	sum := 0
	lines := []string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
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
