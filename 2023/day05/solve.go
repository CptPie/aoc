package day05

import (
	"aoc/2023/utils"
	"errors"
	"fmt"
	"time"
)

func Solve(fileContents []string) ([]string, error) {
	var results []string

	fmt.Println("Solving part 1")
	start := time.Now()
	solution, err := solvePart1(fileContents)
	duration := time.Since(start)

	if err != nil {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 1 Solution: %v\nSolve took: %v\n", solution, duration)

	fmt.Println("Solving part 2")
	start = time.Now()
	solution, err = solvePart2(fileContents)
	duration = time.Since(start)

	if errors.Is(err, utils.NotImplementedError) {
		return results, nil
	} else if err != nil {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 2 Solution: %v\nSolve took: %v\n", solution, duration)
	return results, nil
}

func solvePart1(fileContents []string) (int, error) {
	return 0, utils.NotImplementedError
}

func solvePart2(fileContents []string) (int, error) {
	return 0, utils.NotImplementedError
}
