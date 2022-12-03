# AoC
This repository will eventually contain all my Advent of Code solutions.
It is structured in a way that each year can be solved with a different language since the main "executable" just calls the folder's `entrypoint.sh` which in turn executes the main file of the respective year.

## Usage
Execute the main file `{rootdir}/main.go` i.e. by running `go run main.go {options}`, for all avaliable options see below:

```
Usage: main [--submit] [--download] [--test] YEAR

Positional arguments:
  YEAR                   What years AoC should be solved

Options:
  --submit, -s           Automatically submits solutions for unsolved problems
  --download, -d         Automatically downloads puzzle input if not locally available
  --test, -t             Runs the tests for debugging purposes
  --help, -h             display this help and exit
```
