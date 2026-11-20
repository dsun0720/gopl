package main

import (
	"log"
	"os"
)

func main() {
	if err := reportDuplicates(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
