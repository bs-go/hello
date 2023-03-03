package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func contains(aSlice []string, item string) bool {
	for _, i := range aSlice {
		if i == item {
			return true
		}
	}
	return false
}

func Which() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	files := arguments[1:]

	for _, file := range files {
		execs := make([]string, 0)
		path := os.Getenv("PATH")
		pathSplit := filepath.SplitList(path)

		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					// Is it executable?
					if mode&0111 != 0 {
						if !contains(execs, fullPath) {
							execs = append(execs, fullPath)
						}
					}
				}
			}
		}
		if len(execs) > 0 {
			fmt.Printf("%s: %d Executable(s)\n", file, len(execs))
			fmt.Println(execs)

			// Clear the slice
			execs = nil
		}
	}
}
