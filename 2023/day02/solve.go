package day02

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

type Game struct {
	Id   int
	Sets []GameSet
}

func (g *Game) maxPieces() GameSet {
	maxR, maxG, maxB := 0, 0, 0
	for _, set := range g.Sets {
		if set.Red >= maxR {
			maxR = set.Red
		}
		if set.Green >= maxG {
			maxG = set.Green
		}
		if set.Blue >= maxB {
			maxB = set.Blue
		}
	}
	return GameSet{
		Red:   maxR,
		Green: maxG,
		Blue:  maxB,
	}
}

type GameSet struct {
	Red   int
	Green int
	Blue  int
}

func solvePart1(fileContents []string) (int, error) {

	limit := GameSet{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	sum := 0

	for _, line := range fileContents {

		game, err := constructGame(line)
		if err != nil {
			return 0, err
		}

		maximum := game.maxPieces()
		if !(maximum.Red > limit.Red || maximum.Green > limit.Green || maximum.Blue > limit.Blue) {
			sum += game.Id
		}
	}

	return sum, nil
}

func solvePart2(fileContents []string) (int, error) {
	sum := 0

	for _, line := range fileContents {

		game, err := constructGame(line)
		if err != nil {
			return 0, err
		}

		maximum := game.maxPieces()
		sum += maximum.Green * maximum.Blue * maximum.Red

	}

	return sum, nil

}

func constructGame(line string) (*Game, error) {
	game := Game{}
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("unexpected line format: %s", line)
	}
	sets := parts[1]
	parts = strings.Split(parts[0], " ")
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("number format error: %s", parts[1])
	}
	game.Id = id

	setsSlice := strings.Split(sets, ";")

	for _, set := range setsSlice {
		gameSet := GameSet{}
		set = strings.TrimSpace(set)
		parts = strings.Split(set, ", ")
		for _, part := range parts {
			inner := strings.Split(part, " ")
			if strings.Contains(part, "red") {
				gameSet.Red, err = strconv.Atoi(inner[0])
				if err != nil {
					return nil, fmt.Errorf("number format error: %s", inner[0])
				}
			}
			if strings.Contains(part, "green") {
				gameSet.Green, err = strconv.Atoi(inner[0])
				if err != nil {
					return nil, fmt.Errorf("number format error: %s", inner[0])
				}
			}
			if strings.Contains(part, "blue") {
				gameSet.Blue, err = strconv.Atoi(inner[0])
				if err != nil {
					return nil, fmt.Errorf("number format error: %s", inner[0])
				}
			}
		}
		game.Sets = append(game.Sets, gameSet)
	}
	return &game, nil
}
