package day01

import (
	"2022/utils"
	"fmt"
	"sort"
	"strconv"
)

func Solve() ([]int, error) {
	day := 1

	var results []int

	solution, err := solvePart1(day)

	if err != nil {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(day)

	if err != nil {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil
}

func solvePart1(day int) (int, error) {

	fileScanner, err := utils.ReadFile(day)

	if err != nil {
		return 0, err
	}

	maxSum := 0
	var values []int
	for fileScanner.Scan() {

		if fileScanner.Text() == "" {
			sum := 0
			for _, v := range values {
				sum += v
			}

			if sum > maxSum {
				maxSum = sum
			}

			values = []int{}
		} else {
			value, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				return 0, err
			}
			values = append(values, value)

		}
	}

	return maxSum, nil

}

func solvePart2(day int) (int, error) {
	fileScanner, err := utils.ReadFile(day)

	if err != nil {
		return 0, err
	}

	var sums []int
	var values []int
	for fileScanner.Scan() {

		if fileScanner.Text() == "" {
			sum := 0
			for _, v := range values {
				sum += v
			}

			sums = append(sums, sum)

			values = []int{}
		} else {
			value, err := strconv.Atoi(fileScanner.Text())
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
