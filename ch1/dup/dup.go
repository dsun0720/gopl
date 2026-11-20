package main

import (
	"bufio"
	"fmt"
	"io"
)

func reportDuplicates(r io.Reader, w io.Writer) error {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	for line, n := range counts {
		if n > 1 {
			if _, err := fmt.Fprintf(w, "%d\t%s\n", n, line); err != nil {
				return err
			}
		}
	}
	return nil
}
