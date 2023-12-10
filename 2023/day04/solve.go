package day04

import (
	"aoc/2023/utils"
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

type Card struct {
	Id          int
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
	cards, err := parseInput(fileContents)
	if err != nil {
		return 0, err
	}

	pile := []Card{}
	for _, card := range cards {
		pile = append(pile, card)
	}

	for i, card := range cards {

		toProcess := []Card{}
		for _, pileEntry := range pile {
			if pileEntry.Id == card.Id {
				toProcess = append(toProcess, pileEntry)
			}
		}

		for _, process := range toProcess {
			if process.Matches > 0 {
				for j := 1; j < process.Matches+1; j++ {
					if i+j+1 != len(cards) {
						pile = append(pile, cards[i+j])
					}
				}
			}
		}

	}

	return len(pile), nil
}

func parseInput(lines []string) ([]Card, error) {

	var err error
	cards := []Card{}

	for _, line := range lines {

		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, "  ", " ")

		if line == "" {
			continue
		}

		card := Card{}

		parts := strings.Split(line, ":")

		idStr := strings.Split(parts[0], " ")[1]
		card.Id, err = strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("ID error: %s\n", line)
			return nil, err
		}

		parts = strings.Split(parts[1], "|")
		winning := strings.Split(parts[0], " ")
		numbers := strings.Split(parts[1], " ")

		for _, value := range winning {
			if value == "" {
				continue
			}
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("WNum error: %s\n", line)
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
				fmt.Printf("Num error: %s\n", line)
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
