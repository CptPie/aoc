package day03

import (
	"aoc/2023/utils"
	"errors"
	"fmt"
	"strconv"
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

type Entry struct {
	Type  EntryType
	Line  int
	Start int
	End   int
	Value string
}

func (e Entry) String() string {
	s := ""
	if e.Type == Number {
		s = "N"
	} else {
		s = "S"
	}
	return fmt.Sprintf("[T: %s, L: %d, %d:%d, V: %s]", s, e.Line, e.Start, e.End, e.Value)
}

type EntryType int

const (
	Number EntryType = 0
	Symbol EntryType = 1
	Nil    EntryType = 2
)

type Gear struct {
	Ratio int
}

func solvePart1(fileContents []string) (int, error) {

	entries := generateEntries(fileContents)

	var numbers []Entry
	var symbols []Entry

	for _, entry := range entries {
		if entry.Type == Number {
			numbers = append(numbers, entry)
		} else {
			symbols = append(symbols, entry)
		}
	}

	filteredNumbers := findValidNumbers(numbers, symbols)

	sum := 0

	for _, number := range filteredNumbers {
		num, _ := strconv.Atoi(number.Value)
		sum += num
	}

	return sum, nil
}

func checkNeighbourhood(entry Entry, symbol Entry) bool {
	// check vertical first
	if symbol.Line == entry.Line || symbol.Line == entry.Line+1 || symbol.Line == entry.Line-1 {
		// horizontal next
		hdiff := 999
		if entry.Start <= symbol.End && entry.End >= symbol.Start {
			// it is "within" the bounds
			hdiff = 0
		}
		if entry.Start-1 == symbol.End || entry.End+1 == symbol.Start {
			hdiff = 1
		}

		if hdiff <= 1 {
			return true
		}
	}
	return false
}

func findValidNumbers(numbers []Entry, symbols []Entry) []Entry {
	var filtered []Entry

	for _, entry := range numbers {
		for _, symbol := range symbols {
			if checkNeighbourhood(entry, symbol) {
				filtered = append(filtered, entry)
			}
		}
	}

	return filtered
}

func solvePart2(fileContents []string) (int, error) {
	entries := generateEntries(fileContents)

	var numbers []Entry
	var symbols []Entry

	for _, entry := range entries {
		if entry.Type == Number {
			numbers = append(numbers, entry)
		} else {
			symbols = append(symbols, entry)
		}
	}

	filteredNumbers := findValidNumbers(numbers, symbols)

	solution := findGears(filteredNumbers, symbols)

	return solution, nil
}

func generateEntries(fileContents []string) []Entry {
	entries := []Entry{}

	var lastType EntryType
	num, sym := "", ""
	startIdx := 99999999999
	for i, line := range fileContents {
		for j, char := range line {
			if char == '.' {
				if num != "" || sym != "" {
					var value string
					if lastType == Number {
						value = num
						num = ""
					} else if lastType == Symbol {
						value = sym
						sym = ""
					}
					entries = append(entries, Entry{
						Start: startIdx,
						End:   j - 1,
						Line:  i,
						Value: value,
						Type:  lastType,
					})
				}
				num, sym = "", ""
				startIdx = 99999999999
				lastType = Nil
				continue
			}

			if startIdx > j {
				startIdx = j
			}

			// variable for the current type
			var currType EntryType

			if unicode.IsDigit(char) {
				num += string(char)
				currType = Number
			} else {
				sym += string(char)
				currType = Symbol
			}
			if lastType != currType && lastType != Nil {
				var value string
				if lastType == Number {
					value = num
					num = ""
				} else if lastType == Symbol {
					value = sym
					sym = ""
				}
				entries = append(entries, Entry{
					Start: startIdx,
					End:   j - 1,
					Line:  i,
					Value: value,
					Type:  lastType,
				})
				startIdx = j
			}
			lastType = currType
		}

	}

	return entries
}

func findGears(numbers []Entry, symbols []Entry) int {
	sum := 0

	for _, symbol := range symbols {
		if symbol.Value == "*" {
			nums := []int{}

			for _, number := range numbers {
				if checkNeighbourhood(number, symbol) {
					num, _ := strconv.Atoi(number.Value)
					nums = append(nums, num)
				}
			}

			if len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}
	}

	return sum
}
