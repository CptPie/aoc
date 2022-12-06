package day06

import (
	"errors"
	"fmt"
	"strings"
)

func Solve(fileContents []string) ([]string, error) {
	var results []string

	solution, err := solvePart1(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%v", solution))

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%v", solution))

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

func solvePart1(fileContents []string) (int, error) {
	for _, line := range fileContents {
		return findUniqueSequence(line, 4)
	}
	return 0, nil
}

func hasDuplicates(str string) bool {
	for idx, char := range str {
		strWithout := str[:idx] + str[idx+1:]
		if strings.Contains(strWithout, string(char)) {
			return true
		}
	}
	return false
}

func findUniqueSequence(line string, lenght int) (int, error) {
	for i := lenght; i < len(line); i++ {
		if !hasDuplicates(line[i-lenght : i]) {
			return i, nil
		}
	}
	return 0, errors.New("no sequence code found")
}

func solvePart2(fileContents []string) (int, error) {
	for _, line := range fileContents {
		return findUniqueSequence(line, 14)
	}
	return 0, nil
}
