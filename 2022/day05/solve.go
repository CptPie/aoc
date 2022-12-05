package day05

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Solve(fileContents []string) ([]string, error) {
	var results []string

	solution, err := solvePart1(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, solution)

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

type Stack struct {
	top      int
	Contents []string
}

func (s *Stack) Push(elem string) {
	if s.Contents == nil {
		s.Contents = []string{}
	}
	s.top++
	s.Contents = append(s.Contents, elem)
}

func (s *Stack) Pop() (string, error) {
	if len(s.Contents) == 0 {
		return "", errors.New("Stack empty")
	}
	last := len(s.Contents) - 1
	elem := s.Contents[last]
	s.Contents = s.Contents[:last]
	return elem, nil
}

func (s *Stack) Peek() (string, error) {
	if len(s.Contents) == 0 {
		return "", errors.New("Stack empty")
	}
	return s.Contents[len(s.Contents)-1], nil
}

func solvePart1(fileContents []string) (string, error) {
	var state []string
	var instructions []string
	foundEmptyLine := false

	for _, line := range fileContents {

		if line == "" {
			foundEmptyLine = true
		} else {
			if !foundEmptyLine {
				state = append(state, line)
			} else {
				instructions = append(instructions, line)
			}
		}
	}
	finalState, err := runInstuctionsPart1(parseState(state), instructions)
	if err != nil {
		return "", err
	}

	res := ""
	for i := 1; i < len(finalState)+1; i++ {
		stack := finalState[i]
		top, err := stack.Peek()
		if err != nil {
			return "", err
		}
		res += top

	}

	return res, nil
}

func runInstuctionsPart1(state map[int]Stack, instructions []string) (map[int]Stack, error) {
	for _, line := range instructions {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		src := state[from]
		dst := state[to]
		src, dst, err := moveCrateMover9000(src, dst, amount)
		if err != nil {
			return nil, err
		}
		state[from] = src
		state[to] = dst
	}
	return state, nil
}

func runInstuctionsPart2(state map[int]Stack, instructions []string) (map[int]Stack, error) {
	for _, line := range instructions {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		src := state[from]
		dst := state[to]
		src, dst, err := moveCrateMover9001(src, dst, amount)
		if err != nil {
			return nil, err
		}
		state[from] = src
		state[to] = dst
	}
	return state, nil
}

func moveCrateMover9000(src Stack, dst Stack, amount int) (Stack, Stack, error) {
	for i := 0; i < amount; i++ {
		elem, err := src.Pop()
		if err != nil {
			return Stack{}, Stack{}, err
		}
		dst.Push(elem)
	}
	return src, dst, nil
}

func moveCrateMover9001(src Stack, dst Stack, amount int) (Stack, Stack, error) {

	helper := Stack{}

	for i := 0; i < amount; i++ {
		elem, err := src.Pop()
		if err != nil {
			return Stack{}, Stack{}, err
		}
		helper.Push(elem)
	}

	for i := 0; i < amount; i++ {
		elem, err := helper.Pop()
		if err != nil {
			return Stack{}, Stack{}, err
		}
		dst.Push(elem)
	}
	return src, dst, nil
}

func parseState(state []string) map[int]Stack {
	// Detect how many stacks we need
	maxCols := len(strings.Split(state[len(state)-1], "   "))

	// Prepare map to save the stacks
	res := make(map[int]Stack)
	for i := 1; i <= maxCols; i++ {
		res[i] = Stack{
			top:      0,
			Contents: []string{},
		}
	}

	// Parse the state in reverse beginning from the bottom most line
	for i := len(state) - 2; i >= 0; i-- {

		var cols []string

		// Getting the content of the columns (stacks)
		// Left to right
		for j := 0; j < maxCols; j++ {
			// some magic numbers ain't hurting anyone
			start := j * 4
			end := j*4 + 3

			// will be of the form [*]
			thisCol := state[i][start:end]

			// append to the "results" - but without the [ ]
			cols = append(cols, string(thisCol[1]))

			// if the end (']') is equal to the length of the line its the last entry
			if len(state[i]) == end {
				break
			}

		}

		// we only care for nonempty stack entries
		for col, val := range cols {
			if val != " " {
				stack := res[col+1]
				stack.Push(val)
				res[col+1] = stack
			}
		}

	}

	return res
}
func solvePart2(fileContents []string) (string, error) {
	var state []string
	var instructions []string
	foundEmptyLine := false

	for _, line := range fileContents {

		if line == "" {
			foundEmptyLine = true
		} else {
			if !foundEmptyLine {
				state = append(state, line)
			} else {
				instructions = append(instructions, line)
			}
		}
	}
	finalState, err := runInstuctionsPart2(parseState(state), instructions)
	if err != nil {
		return "", err
	}

	res := ""
	for i := 1; i < len(finalState)+1; i++ {
		stack := finalState[i]
		top, err := stack.Peek()
		if err != nil {
			return "", err
		}
		res += top

	}

	return res, nil
}
