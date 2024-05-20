/*
This version will read the entire input into memory
in one big gulp, split it into lines all at once, then
process the lines.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile reads the entire contents of a
		// named file.
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// strings.Split splits a string into a slice of
		// substrings based on a separator, in this case
		// a newline character "\n".
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
