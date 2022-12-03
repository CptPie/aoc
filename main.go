package main

import (
	"fmt"
	"os/exec"

	"github.com/alexflint/go-arg"
)

type args struct {
	Year     int  `arg:"positional, required" help:"What years AoC should be solved"`
	Submit   bool `arg:"-s, --submit" help:"Automatically submits solutions for unsolved problems"`
	Download bool `arg:"-d, --download" help:"Automatically downloads puzzle input if not locally available"`
	Test     bool `arg:"-t, --test" help:"Runs the tests for debugging purposes"`
}

func main() {
	var args args

	arg.MustParse(&args)

	dl := ""
	if args.Download {
		dl = "-d"
	}

	submit := ""
	if args.Submit {
		submit = "-s"
	}

	test := ""
	if args.Test {
		test = "-t"
	}

	out, err := exec.Command(
		"/bin/bash",
		fmt.Sprintf("./%v/entrypoint.sh", args.Year),
		dl,
		submit,
		test).Output()
	if err != nil {
		fmt.Printf("[Error]: %v", err)
	}
	fmt.Printf("%s\n", out)
}
