// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		// When a map is passed to a function, the function receives
		// a copy of the reference to the map, so any changes the called
		// function makes to the data structure will be visible through
		// the caller's map reference too.
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplicate2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%s occurs in files: ", line)
			for filename := range files {
				fmt.Printf("%s ", filename)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]bool)
		}
		counts[line][filename] = true
	}
}
