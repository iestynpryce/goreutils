package main

/* Copy file a to b */

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func check_and_error(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", e)
		os.Exit(-1)
	}
}

func main() {
	flag.Parse()

	args := flag.Args()
	var nargs int = len(args)
	if nargs > 2 {
		fmt.Fprintf(os.Stderr, "Multiple source files not supported\n")
		os.Exit(-1)
	}
	if nargs < 2 {
		fmt.Fprintf(os.Stderr, "You must supply a target\n")
		os.Exit(-1)
	}

	src, err := os.Open(args[0])
	check_and_error(err)

	dst, err := os.Create(args[1])
	check_and_error(err)

	_, err = io.Copy(dst, src)
	check_and_error(err)
}
