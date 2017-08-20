package main

import (
	"flag"
	"fmt"
	"os"
)

func printError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", e)
	}
}

func main() {
	parents := flag.Bool("p", false, "create directory parents")

	flag.Parse()

	args := flag.Args()

	for _, f := range args {
		if *parents {
			err := os.MkdirAll(f, os.ModeDir|0777)
			printError(err)
		} else {
			err := os.Mkdir(f, os.ModeDir|0777)
			printError(err)

		}
	}
}
