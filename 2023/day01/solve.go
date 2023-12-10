package day01

import (
	"aoc/2023/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	sum := 0
	for _, line := range fileContents {
		var digits []string
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
			}
		}

		valStr := ""
		if len(digits) > 1 {
			valStr = fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1])
		} else if len(digits) == 1 {
			valStr = fmt.Sprintf("%s%s", digits[0], digits[0])
		} else {
			valStr = "0"
		}

		val, err := strconv.Atoi(valStr)
		if err != nil {
			return 0, fmt.Errorf("invalid number format %v", digits)
		}
		sum += val
	}
	return sum, nil
}

func solvePart2(fileContents []string) (int, error) {
	sum := 0

	mapOfNumbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range fileContents {
		var digits []string
		for idx, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
			} else {
				for key, value := range mapOfNumbers {
					if strings.HasPrefix(line[idx:], key) {
						digits = append(digits, fmt.Sprintf("%d", value))
					}
				}
			}
		}

		valStr := ""
		if len(digits) > 1 {
			valStr = fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1])
		} else {
			valStr = fmt.Sprintf("%s%s", digits[0], digits[0])
		}

		val, err := strconv.Atoi(valStr)
		if err != nil {
			return 0, fmt.Errorf("invalid number format %v", digits)
		}
		sum += val
	}
	return sum, nil

}
