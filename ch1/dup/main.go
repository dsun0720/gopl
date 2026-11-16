package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v\t%s\n", n, line)
		}
	}
}
