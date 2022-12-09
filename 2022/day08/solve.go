package day08

import (
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

	results = append(results, fmt.Sprintf("%v", solution))

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution, err = solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%v", solution))

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

type Tree struct {
	IsVisible  bool
	Height     int
	SightUp    int
	SightDown  int
	SightLeft  int
	SightRight int
}

func (t Tree) Score() int {
	return t.SightUp * t.SightDown * t.SightLeft * t.SightRight
}

func (t Tree) String() string {
	return fmt.Sprintf("[U %d, D %d, L %d, R %d, S %d]", t.SightUp, t.SightDown, t.SightLeft, t.SightRight, t.Score())
}

func solvePart1(fileContents []string) (int, error) {

	treeMap, err := processInput(fileContents)
	if err != nil {
		return 0, err
	}

	treeMap, err = analyzeTreeMapPart1(treeMap)

	if err != nil {
		return 0, err
	}

	amount := 0
	for _, trees := range treeMap {
		for _, tree := range trees {
			if tree.IsVisible {
				amount += 1
			}
		}
	}

	return amount, nil
}

func analyzeTreeMapPart1(treeMap [][]Tree) ([][]Tree, error) {
	for i := 1; i < len(treeMap)-1; i++ {
		inner := treeMap[i][1 : len(treeMap[i])-1]
		for j, tree := range inner {

			if isVisible(treeMap, tree, i, j+1) {
				tree.IsVisible = true
				treeMap[i][j+1] = tree
			}
		}
	}
	return treeMap, nil
}

func isVisible(treeMap [][]Tree, tree Tree, row int, column int) bool {
	// Assume it will not be visible (proof by contradiction)
	// If we find one direction where it is visible, we can return

	// check above
	visibiltyUp := true
	for i := 0; i < row; i++ {
		if treeMap[i][column].Height >= tree.Height {
			visibiltyUp = false
		}
	}

	// check below
	visibiltyDown := true
	for i := len(treeMap) - 1; i > row; i-- {
		if treeMap[i][column].Height >= tree.Height {
			visibiltyDown = false
		}
	}

	// check left
	visibiltyLeft := true
	for i := 0; i < column; i++ {
		if treeMap[row][i].Height >= tree.Height {
			visibiltyLeft = false
		}
	}

	// check right
	visibiltyRight := true
	for i := len(treeMap[row]) - 1; i > column; i-- {
		if treeMap[row][i].Height >= tree.Height {
			visibiltyRight = false
		}
	}

	return visibiltyUp || visibiltyDown || visibiltyLeft || visibiltyRight
}

func processInput(contents []string) ([][]Tree, error) {
	var rows [][]Tree
	for i, line := range contents {
		parts := strings.Split(line, "")

		var row []Tree
		for j, entry := range parts {
			height, err := strconv.Atoi(entry)
			if err != nil {
				return nil, err
			}
			row = append(row, Tree{
				IsVisible: i == 0 || i == len(contents)-1 || j == 0 || j == len(parts)-1,
				Height:    height,
			})
		}

		rows = append(rows, row)
	}
	return rows, nil
}
func solvePart2(fileContents []string) (int, error) {
	treeMap, err := processInput(fileContents)
	if err != nil {
		return 0, err
	}

	treeMap, err = analyzeTreeMapPart2(treeMap)

	if err != nil {
		return 0, err
	}

	maxScore := 0
	for _, trees := range treeMap {
		for _, tree := range trees {
			if tree.Score() > maxScore {
				maxScore = tree.Score()
			}
		}
	}

	return maxScore, nil
}

func analyzeTreeMapPart2(treeMap [][]Tree) ([][]Tree, error) {
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[i])-1; j++ {
			tree := analyzeTree(i, j, treeMap)
			treeMap[i][j] = tree
		}
	}
	return treeMap, nil
}

func analyzeTree(row int, column int, treeMap [][]Tree) Tree {
	tree := treeMap[row][column]

	// check above
	visibiltyUp := 0
	for i := row - 1; i >= 0; i-- {
		other := treeMap[i][column]
		if other.Height >= tree.Height {
			visibiltyUp += 1
			break
		} else {
			visibiltyUp += 1
		}
	}

	// check below
	visibiltyDown := 0
	for i := row + 1; i <= len(treeMap)-1; i++ {
		other := treeMap[i][column]
		if other.Height >= tree.Height {
			visibiltyDown += 1
			break
		} else {
			visibiltyDown += 1
		}
	}

	// check left
	visibiltyLeft := 0
	for i := column - 1; i >= 0; i-- {
		other := treeMap[row][i]
		if other.Height >= tree.Height {
			visibiltyLeft += 1
			break
		} else {
			visibiltyLeft += 1
		}
	}

	// check right
	visibiltyRight := 0
	for i := column + 1; i <= len(treeMap[row])-1; i++ {
		other := treeMap[row][i]
		if other.Height >= tree.Height {
			visibiltyRight += 1
			break
		} else {
			visibiltyRight += 1
		}
	}

	return Tree{
		IsVisible:  tree.IsVisible,
		Height:     tree.Height,
		SightUp:    visibiltyUp,
		SightDown:  visibiltyDown,
		SightLeft:  visibiltyLeft,
		SightRight: visibiltyRight,
	}
}
