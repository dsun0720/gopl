package main

import (
	"fmt"
	"io"
)

func echoCounted(w io.Writer, args []string) error {
	var s string
	for i := 0; i < len(args); i++ {
		if i > 0 {
			s += " "
		}
		s += args[i]
	}
	return printEcho(w, s)
}

func echoCondition(w io.Writer, args []string) error {
	var s string
	i := 0
	for i < len(args) {
		if i > 0 {
			s += " "
		}
		s += args[i]
		i++
	}
	return printEcho(w, s)
}

func echoRange(w io.Writer, args []string) error {
	var s string
	for i, arg := range args {
		if i > 0 {
			s += " "
		}
		s += arg
	}
	return printEcho(w, s)
}

func echoBreak(w io.Writer, args []string) error {
	var s string
	i := 0
	for {
		if i == len(args) {
			break
		}
		if i > 0 {
			s += " "
		}
		s += args[i]
		i++
	}
	return printEcho(w, s)
}

func printEcho(w io.Writer, s string) error {
	_, err := fmt.Fprintln(w, s)
	return err
}
