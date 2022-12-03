package day01

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve(fileContents []string) ([]int, error) {

	var results []int

	solution, err := solvePart1(fileContents)

	if err != nil {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if err != nil {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil
}

func solvePart1(fileContents []string) (int, error) {

	maxSum := 0
	var values []int
	for _, line := range fileContents {

		if line == "" {
			sum := 0
			for _, v := range values {
				sum += v
			}

			if sum > maxSum {
				maxSum = sum
			}

			values = []int{}
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}
			values = append(values, value)

		}
	}

	return maxSum, nil

}

func solvePart2(fileContents []string) (int, error) {

	var sums []int
	var values []int
	for _, line := range fileContents {

		if line == "" {
			sum := 0
			for _, v := range values {
				sum += v
			}

			sums = append(sums, sum)

			values = []int{}
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}
			values = append(values, value)

		}
	}

	sort.Ints(sums)
	sum := 0
	for _, val := range sums[len(sums)-3:] {
		sum += val
	}
	return sum, nil
}
