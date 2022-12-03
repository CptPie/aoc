package main

import (
	"2022/day01"
	"2022/day02"
	"2022/day03"
	"2022/utils"
	"errors"
	"fmt"
	"github.com/alexflint/go-arg"
)

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
			fmt.Println("Downloading input file")
			err := utils.GetDayInput(i, args.ConfigPath)
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
			err = utils.SubmitSolutions(i, results, args.ConfigPath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func solve(day int) ([]int, error) {
	var err error
	switch day {
	case 1:
		contents, err := utils.ReadFile(day)
		if err != nil {
			return nil, err
		}
		return day01.Solve(contents)
	case 2:
		contents, err := utils.ReadFile(day)
		if err != nil {
			return nil, err
		}
		return day02.Solve(contents)
	case 3:
		contents, err := utils.ReadFile(day)
		if err != nil {
			return nil, err
		}
		return day03.Solve(contents)
	case 4:
		err = errors.New("solve for day 4 not implemented")
	case 5:
		err = errors.New("solve for day 5 not implemented")
	case 6:
		err = errors.New("solve for day 6 not implemented")
	case 7:
		err = errors.New("solve for day 7 not implemented")
	case 8:
		err = errors.New("solve for day 8 not implemented")
	case 9:
		err = errors.New("solve for day 9 not implemented")
	case 10:
		err = errors.New("solve for day 10 not implemented")
	case 11:
		err = errors.New("solve for day 11 not implemented")
	case 12:
		err = errors.New("solve for day 12 not implemented")
	case 13:
		err = errors.New("solve for day 13 not implemented")
	case 14:
		err = errors.New("solve for day 14 not implemented")
	case 15:
		err = errors.New("solve for day 15 not implemented")
	case 16:
		err = errors.New("solve for day 16 not implemented")
	case 17:
		err = errors.New("solve for day 17 not implemented")
	case 18:
		err = errors.New("solve for day 18 not implemented")
	case 19:
		err = errors.New("solve for day 19 not implemented")
	case 20:
		err = errors.New("solve for day 20 not implemented")
	case 21:
		err = errors.New("solve for day 21 not implemented")
	case 22:
		err = errors.New("solve for day 22 not implemented")
	case 23:
		err = errors.New("solve for day 23 not implemented")
	case 24:
		err = errors.New("solve for day 24 not implemented")
	case 25:
		err = errors.New("solve for day 25 not implemented")
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
