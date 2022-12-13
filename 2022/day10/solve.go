package day10

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

	results = append(results, fmt.Sprintf("%v", solution))

	fmt.Printf("Part 1 Solution: %v\n", solution)

	solution2, err := solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%v", solution2))

	fmt.Printf("Part 2 Solution: %v\n", solution2)
	return results, nil

}

type CPU struct {
	Registers map[string]Register
	PC        int
}

var CRTPixels [240]string

func (c *CPU) Tick() {

	// X register value is pointer to the middle of the 3 wide drawing pointer
	// Do CRT Tick
	if c.PC%40 == c.Registers["X"].Value {
		CRTPixels[c.PC] = "#"
	} else if c.PC%40 == c.Registers["X"].Value-1 {
		CRTPixels[c.PC] = "#"
	} else if c.PC%40 == c.Registers["X"].Value+1 {
		CRTPixels[c.PC] = "#"
	} else {
		CRTPixels[c.PC] = "."
	}
	c.PC++
	if c.PC == 20 || (c.PC-20)%40 == 0 {
		Cycles[c.PC] = c.PC * c.Registers["X"].Value
	}
	//fmt.Printf("PC: %d, Pos: %d\n", c.PC, c.Registers["X"].Value)
	//printCRT()
}

type Register struct {
	Name  string
	Value int
}

type Instruction interface {
	Run(cpu *CPU)
}

type AddxInstruction struct {
	Value int
}

func (a AddxInstruction) Run(cpu *CPU) {
	cpu.Tick()
	register := cpu.Registers["X"]
	register.Value += a.Value
	cpu.Tick()
	cpu.Registers["X"] = register
	return
}

type NoopInstruction struct {
}

func (n NoopInstruction) Run(cpu *CPU) {
	cpu.Tick()
	return
}

var Cycles = map[int]int{}

func solvePart1(fileContents []string) (int, error) {
	cpu := CPU{
		Registers: make(map[string]Register),
		PC:        0,
	}

	cpu.Registers["X"] = Register{
		Name:  "X",
		Value: 1,
	}

	for _, line := range fileContents {
		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			return 0, errors.New("Could not parse line")
		}
		switch parts[0] {
		case "noop":
			NoopInstruction{}.Run(&cpu)
		case "addx":
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				return 0, err
			}
			AddxInstruction{Value: val}.Run(&cpu)
		}
	}

	sum := 0

	for _, val := range Cycles {
		sum += val
	}
	return sum, nil
}

func solvePart2(fileContents []string) (string, error) {
	cpu := CPU{
		Registers: make(map[string]Register),
		PC:        0,
	}

	cpu.Registers["X"] = Register{
		Name:  "X",
		Value: 1,
	}

	// simulate cpu
	for _, line := range fileContents {
		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			return "", errors.New("Could not parse line")
		}
		switch parts[0] {
		case "noop":
			NoopInstruction{}.Run(&cpu)
		case "addx":
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
			AddxInstruction{Value: val}.Run(&cpu)
		}
	}

	return printCRT(), nil
}

func printCRT() string {
	str := ""
	for i, pixel := range CRTPixels {
		str += pixel
		if (i+1)%40 == 0 {
			str += "\n"
		}
	}
	return str
}
