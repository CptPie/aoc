package day07

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

	solution, err = solvePart2(fileContents)

	if err != nil && err.Error() != "Not implemented" {
		return nil, err
	}

	results = append(results, fmt.Sprintf("%v", solution))

	fmt.Printf("Part 2 Solution: %v\n", solution)
	return results, nil

}

type Directory struct {
	Name        string
	Files       []File
	Directories []*Directory
	IsRoot      bool
	Parent      *Directory
	Level       int
}

func (d Directory) Path() string {
	if d.IsRoot {
		return "/"
	} else {
		return d.Parent.Path() + "/" + d.Name
	}
}

func (d Directory) Size() int {
	sum := 0
	for _, file := range d.Files {
		sum += file.Size
	}
	for _, dir := range d.Directories {
		sum += dir.Size()
	}
	return sum
}

func (d Directory) Print() {
	fmt.Printf("%s%8d %s/\n", strings.Repeat("   |", d.Level), d.Size(), d.Name)
	for _, file := range d.Files {
		fmt.Printf("%s%8d - %s\t\n", strings.Repeat("   |", d.Level+1), file.Size, file.Name)
	}
	for _, dir := range d.Directories {
		dir.Print()
	}
}

type File struct {
	Name string
	Size int
}

func solvePart1(fileContents []string) (int, error) {

	dirs, err := buildFileSystem(fileContents)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, dir := range dirs {
		if dir.Size() <= 100000 {
			//fmt.Printf("Dir: %s, Size: %d\n",dir.Name,dir.Size())
			sum += dir.Size()
		}
	}

	return sum, nil

}

func buildFileSystem(fileContents []string) (map[string]Directory, error) {

	dirs := make(map[string]Directory)

	var currentDir *Directory

	rootDir := Directory{
		Name:        "/",
		Files:       nil,
		Directories: nil,
		IsRoot:      true,
		Parent:      nil,
		Level:       0,
	}

	dirs["/"] = rootDir

	for _, line := range fileContents {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			//command
			switch parts[1] {
			case "cd":
				dirName := parts[2]
				switch dirName {
				case "..":
					// change dir to parent
					if !currentDir.IsRoot {
						nextDir := dirs[currentDir.Parent.Path()]
						currentDir = &nextDir
						//currentDir = currentDir.Parent
					}
				case "/":
					// change dir to root dir
					//currentDir = &rootDir
					nextDir := dirs["/"]
					currentDir = &nextDir
				default:
					// change dir to parts[2]
					path := currentDir.Path() + "/" + parts[2]
					nextDir, ok := dirs[path]
					if ok {
						currentDir = &nextDir
					} else {
						nextDir = Directory{
							Name:        parts[2],
							Files:       nil,
							Directories: nil,
							IsRoot:      false,
							Parent:      currentDir,
							Level:       currentDir.Level + 1,
						}
						dirs[path] = nextDir
						currentDir = &nextDir
					}
				}
			case "ls":

			default:
				return nil, errors.New(fmt.Sprintf("unknown command: %s", line))
			}
		} else {
			if parts[0] == "dir" {
				// directory listing
				path := currentDir.Path() + "/" + parts[1]
				newDir := Directory{
					Name:        parts[1],
					Files:       nil,
					Directories: nil,
					IsRoot:      parts[1] == "/",
					Parent:      currentDir,
					Level:       currentDir.Level + 1,
				}
				dirs[path] = newDir
				currentDir.Directories = append(currentDir.Directories, &newDir)
				dirs[currentDir.Path()] = *currentDir
			} else {
				// file listing
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					return nil, err
				}
				name := parts[1]
				file := File{
					Name: name,
					Size: size,
				}
				currentDir.Files = append(currentDir.Files, file)
				dirs[currentDir.Path()] = *currentDir
			}
		}
	}

	maxLevel := 0
	for _, dir := range dirs {
		if dir.Level > maxLevel {
			maxLevel = dir.Level
		}
	}

	// update subdirs, depth first
	for level := maxLevel; level >= 0; level-- {
		for _, dir := range dirs {
			if dir.Level == level {
				//fmt.Printf("Updating dirs of level %d\n",level)
				if dir.Directories != nil && len(dir.Directories) != 0 {
					// has subdirs
					for i, directory := range dir.Directories {
						// update subdir
						updatedDir := dirs[directory.Path()]
						dir.Directories[i] = &updatedDir
					}
				}
			}
		}
	}
	return dirs, nil
}

func solvePart2(fileContents []string) (int, error) {
	dirs, err := buildFileSystem(fileContents)
	if err != nil {
		return 0, err
	}

	const fileSystemTotal = 70000000

	totalUsed := dirs["/"].Size()

	totalFree := fileSystemTotal - totalUsed

	const required = 30000000

	toDelete := required - totalFree

	fmt.Printf("Used: %d/%d\n",totalUsed,fileSystemTotal)
	fmt.Printf("We need %d, free is %d so we need to delete %d\n",required,totalFree,toDelete)

	var candidates []Directory
	for _, dir := range dirs {
		if dir.Size() >= toDelete {
			candidates = append(candidates, dir)
		}
	}

	dummyFile := File{
		Name: "",
		Size: fileSystemTotal,
	}

	smallest := Directory{
		Name:        "",
		Files:       []File{
			dummyFile,
		},
		Directories: nil,
		IsRoot:      false,
		Parent:      nil,
		Level:       0,
	}

	for _, candidate := range candidates {
		if candidate.Size() < smallest.Size() {
			smallest = candidate
		}
	}

	return smallest.Size(), nil
}
