package day02

import (
	"aoc/2025/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

	if len(fileContents) > 1 {
		return 0, fmt.Errorf("Unexpected lenght of the input file, expected only one line, got %d", len(fileContents))
	}

	parts := strings.Split(fileContents[0], ",")

	sum := 0
	for _, part := range parts {
		nums := strings.Split(part, "-")
		if len(nums) != 2 {
			return 0, fmt.Errorf("shit's fucked bro")
		}

		a, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, fmt.Errorf("Could not convert %s to int", nums[0])
		}

		b, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, fmt.Errorf("Could not convert %s to int", nums[0])
		}

		for i := a; i <= b; i++ {
			iStr := strconv.Itoa(i)
			if (len(iStr) % 2) == 0 {
				half1 := iStr[:len(iStr)/2]
				half2 := iStr[len(iStr)/2:]

				if half1 == half2 {
					sum += i
				}
			}
		}

	}

	return sum, nil
}

func solvePart2(fileContents []string) (int, error) {
	if len(fileContents) > 1 {
		return 0, fmt.Errorf("Unexpected lenght of the input file, expected only one line, got %d", len(fileContents))
	}

	parts := strings.Split(fileContents[0], ",")

	sum := 0
	for _, part := range parts {
		nums := strings.Split(part, "-")
		if len(nums) != 2 {
			return 0, fmt.Errorf("shit's fucked bro")
		}

		a, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, fmt.Errorf("Could not convert %s to int", nums[0])
		}

		b, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, fmt.Errorf("Could not convert %s to int", nums[0])
		}

		for i := a; i <= b; i++ {
			iStr := strconv.Itoa(i)

			maxSegments := len(iStr)

			for lenght := 1; lenght <= maxSegments; lenght++ {
				segments, err := splitString(iStr, lenght)
				if err != nil {
					continue
				}
				if len(segments) <= 1 {
					continue
				}

				checkSegment := segments[0]
				fine := true
				for _, segment := range segments {
					if checkSegment != segment {
						fine = false
						break
					}
				}

				if fine {
					sum += i
					break
				}
			}
		}
	}

	return sum, nil
}

func splitString(str string, size int) ([]string, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size must be positive, got %d", size)
	}

	if len(str)%size != 0 {
		return nil, fmt.Errorf("string is not divisible by size")
	}

	var result []string

	for i := 0; i < len(str); i += size {
		result = append(result, str[i:i+size])
	}
	return result, nil
}
