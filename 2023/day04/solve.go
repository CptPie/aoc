package day04

import (
	"aoc/2023/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Solve(fileContents []string) ([]string, error) {

	var results []string

	solution, err := solvePart1(fileContents)

	if err != nil {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if errors.Is(err, utils.NotImplementedError) {
		return results, nil
	} else if err != nil {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil
}

type Card struct {
	WinningNums []int
	Nums        []int
	Matches     int
	Worth       int
}

func solvePart1(fileContents []string) (int, error) {
	cards, err := parseInput(fileContents)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, card := range cards {
		sum += card.Worth
	}

	return sum, nil
}

func solvePart2(fileContents []string) (int, error) {
	return 0, utils.NotImplementedError
}

func parseInput(lines []string) ([]Card, error) {
	cards := []Card{}

	for _, line := range lines {

		if line == "" {
			continue
		}

		card := Card{}
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], "|")
		winning := strings.Split(parts[0], " ")
		numbers := strings.Split(parts[1], " ")

		for _, value := range winning {
			if value == "" {
				continue
			}
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			card.WinningNums = append(card.WinningNums, num)
		}

		for _, value := range numbers {
			if value == "" {
				continue
			}
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			card.Nums = append(card.Nums, num)
		}

		worth := 0
		matches := false
		for _, win := range card.WinningNums {
			for _, num := range card.Nums {
				if win == num {
					card.Matches++
					if !matches {
						matches = true
						worth = 1
					} else {
						worth = worth * 2
					}
				}
			}
		}

		card.Worth = worth

		cards = append(cards, card)
	}

	return cards, nil
}
