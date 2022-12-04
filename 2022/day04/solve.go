package day04

import (
	"errors"
	"fmt"
)

func Solve(fileContents []string) ([]int, error) {
	var results []int

	solution, err := solvePart1(fileContents)

	if err != nil && err.Error() != "Not implemented" {
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
	return 0, errors.New("Not implemented")
}
func solvePart2(fileContents []string) (int, error) {
	return 0, errors.New("Not implemented")
}
