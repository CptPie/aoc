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
		var amount int
		var err error
		oldPos := arrowPos

		if line[0] == 'R' {
			amount, err = strconv.Atoi(line[1:])
			if err != nil {
				return 0, err
			}

			newPos := oldPos + amount
			arrowPos = newPos % 100

			// Count crossings DURING rotation (not including final position if it's 0)
			crossings := 0
			if newPos%100 == 0 && newPos > 0 {
				// We land on 0, so don't count it as a crossing
				crossings = (newPos / 100) - 1
			} else {
				crossings = newPos / 100
			}
			count += crossings

			// If we land on 0, count it separately
			if arrowPos == 0 {
				count++
			}

		} else if line[0] == 'L' {
			amount, err = strconv.Atoi(line[1:])
			if err != nil {
				return 0, err
			}

			// Count crossings going left
			// Special case: if starting from 0 and going left, we don't cross 0 initially
			if oldPos == 0 {
				// Starting from 0: only count full rotations beyond the start
				if amount >= 100 {
					crossings := amount / 100
					// But if we land back on 0, don't count it as a crossing
					if amount%100 == 0 {
						crossings--
					}
					count += crossings
				}
			} else if amount > oldPos {
				// We cross 0 at least once
				crossings := 0
				if (amount-oldPos)%100 == 0 {
					// We land on 0, so don't count it as a crossing
					crossings = (amount - oldPos) / 100
				} else {
					crossings = ((amount - oldPos - 1) / 100) + 1
				}
				count += crossings
			}

			arrowPos = ((oldPos-amount)%100 + 100) % 100

			// If we land on 0, count it separately
			if arrowPos == 0 {
				count++
			}

		} else {
			return 0, fmt.Errorf("Invalid line format, expected to start with L or R, got %s", line)
		}
	}

	return count, nil
}
