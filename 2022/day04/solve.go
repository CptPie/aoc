package day04

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func Solve(fileContents []string) ([]string, error) {
	var results []string

	solution, err := solvePart1(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%d", solution))

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

func solvePart1(fileContents []string) (int, error) {
	cnt := 0
	for _, line := range fileContents {
		listA, listB, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		if sliceFullyContainsSlice(listA, listB) {
			cnt++
		}
	}
	return cnt, nil
}

func parseLine(line string) ([]int, []int, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return nil, nil, errors.New("Error while parsing line")
	}
	var sliceA []int
	var sliceB []int
	for i, part := range parts {
		idxs := strings.Split(part, "-")
		if len(idxs) != 2 {
			return nil, nil, errors.New("Error while parsing line")
		}
		beginning, err := strconv.Atoi(idxs[0])
		if err != nil {
			return nil, nil, err
		}
		end, err := strconv.Atoi(idxs[1])
		if err != nil {
			return nil, nil, err
		}
		for j := beginning; j <= end; j++ {
			switch i {
			case 0:
				sliceA = append(sliceA, j)
			case 1:
				sliceB = append(sliceB, j)
			}
		}
	}
	return sliceA, sliceB, nil
}

func sliceFullyContainsSlice(A []int, B []int) bool {
	if len(A) > len(B) {
		for _, e := range B {
			if !slices.Contains(A, e) {
				return false
			}
		}
		return true

	} else {
		for _, e := range A {
			if !slices.Contains(B, e) {
				return false
			}
		}
		return true
	}
}

func solvePart2(fileContents []string) (int, error) {
	cnt := 0
	for _, line := range fileContents {
		listA, listB, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		if sliceContainsContent(listA, listB) {
			cnt++
		}
	}
	return cnt, nil
}

func sliceContainsContent(A []int, B []int) bool {
	for _, e := range A {
		if slices.Contains(B, e) {
			return true
		}
	}
	return false
}
