package day01

import (
	"aoc/2023/utils"
	"errors"
	"fmt"
	"math"
	"strconv"
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
	count := 0
	arrowPos := 50

	for _, line := range fileContents {
		if line[0] == 'R' {
			// turning the dial right, in fact is just addition
			amount, err := strconv.Atoi(line[1:])
			if err != nil {
				return 0, err
			}

			amount = amount % 100

			//			fmt.Printf("Old pos: %d, turning %c %d times, ", arrowPos, line[0], amount)

			arrowPos += amount

			// add wrap around, once we pass 100, roll back to 0
			if arrowPos >= 100 {
				arrowPos -= 100
			}

			//			fmt.Printf("new pos: %d\n", arrowPos)

		} else if line[0] == 'L' {
			// turning the dial left, in fact is just subtraction
			amount, err := strconv.Atoi(line[1:])
			if err != nil {
				return 0, err
			}

			amount = amount % 100
			//						fmt.Printf("Old pos: %d, turning %c %d times, ", arrowPos, line[0], amount)
			arrowPos -= amount

			// wrap around - if we end up with a negative number, adjust the value (99-x)
			if arrowPos < 0 {
				arrowPos = 100 - int(math.Abs(float64(arrowPos)))
			}

			//			fmt.Printf("new pos: %d\n", arrowPos)
		} else {
			return 0, fmt.Errorf("Invalid line format, expected to start with L or R, got %s", line)
		}

		if arrowPos == 0 {
			//			fmt.Println("Arrow at 0")
			count++
		}

	}

	return count, nil
	// return 0, utils.NotImplementedError
}

func solvePart2(fileContents []string) (int, error) {
	count := 0

	arrowPos := 50

	for _, line := range fileContents {
		zeroCrosses := 0
		lastPos := arrowPos
		if line[0] == 'R' {
			// turning the dial right, in fact is just addition
			amount, err := strconv.Atoi(line[1:])
			if err != nil {
				return 0, err
			}

			zeroCrosses = amount / 100
			amount = amount % 100

			//			fmt.Printf("Old pos: %d, turning %c %d times, ", arrowPos, line[0], amount)

			arrowPos += amount

			// add wrap around, once we pass 100, roll back to 0
			if arrowPos >= 100 {
				if lastPos != 0 {
					zeroCrosses++
				}
				arrowPos -= 100
			}

			//			fmt.Printf("new pos: %d\n", arrowPos)

		} else if line[0] == 'L' {
			// turning the dial left, in fact is just subtraction
			amount, err := strconv.Atoi(line[1:])
			if err != nil {
				return 0, err
			}

			zeroCrosses = amount / 100
			amount = amount % 100
			//			fmt.Printf("Old pos: %d, turning %c %d times, ", arrowPos, line[0], amount)
			arrowPos -= amount

			// wrap around - if we end up with a negative number, adjust the value (99-x)
			if arrowPos < 0 {
				if lastPos != 0 {
					zeroCrosses++
				}
				arrowPos = 100 - int(math.Abs(float64(arrowPos)))
			}

			//			fmt.Printf("new pos: %d", arrowPos)
		} else {
			return 0, fmt.Errorf("Invalid line format, expected to start with L or R, got %s", line)
		}

		if arrowPos == 0 {
			//			fmt.Println("Arrow at 0")
			count++
		} else {
			//			fmt.Printf(" found %d zeroCrosses\n", zeroCrosses)
			count += zeroCrosses
		}

	}

	return count, nil
}
