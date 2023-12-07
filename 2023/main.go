package main

import (
	"aoc/2023/day01"
	"aoc/2023/day02"
	"aoc/2023/utils"
	"errors"
	"fmt"

	"github.com/alexflint/go-arg"
)

const year = 2023

type args struct {
	Submit     bool   `arg:"-s, --submit" help:"Submit calculated results of the provided day"`
	ConfigPath string `arg:"-c, --config" default:"config.json" help:"Path to the config file. Defaults to config.json"`
	Download   bool   `arg:"-d,--download" help:"Downloads the puzzle input for [DAY]. If no day is provided it will download all available puzzle inputs"`
	Test       bool   `arg:"-t, --test"`
}

func main() {
	var args args

	arg.MustParse(&args)

	run(args)
}

func run(args args) {

	for i := 1; i < 25; i++ {
		fmt.Printf("----- Day %d -----\n", i)
		if args.Download {
			fmt.Println("Downloading puzzle prompt")
			err := utils.GetDayDesc(year, i, args.ConfigPath)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Downloading input file")
			err = utils.GetDayInput(year, i, args.ConfigPath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Println("Solving")
		results, err := solve(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		if args.Submit {
			fmt.Println("Submitting results")
			err = utils.SubmitSolutions(year, i, results, args.ConfigPath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func solve(day int) ([]string, error) {
	var err error
	contents, err := utils.ReadFile(fmt.Sprintf("day%02d/input", day))

	if err != nil {
		return nil, err
	}
	switch day {
	case 1:
		return day01.Solve(contents)
	case 2:
		return day02.Solve(contents)
	case 3:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 4:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 5:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 6:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 7:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 8:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 9:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 10:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 11:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 12:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 13:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 14:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 15:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 16:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 17:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 18:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 19:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 20:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 21:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 22:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 23:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 24:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	case 25:
		err = errors.New(fmt.Sprintf("solve for day %d not implemented", day))
	default:
		if day < 0 || day > 25 {
			return nil, errors.New("day out of bounds")
		}
	}
	if err != nil {
		return nil, err
	}
	return nil, nil
}
