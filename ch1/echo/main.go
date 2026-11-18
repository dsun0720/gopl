package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, echo := range []func(io.Writer, []string) error{
		echoCounted, echoCondition, echoRange, echoBreak,
	} {
		if err := echo(os.Stdout, args); err != nil {
			log.Fatal(err)
		}
	}
}
