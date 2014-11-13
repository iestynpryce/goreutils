package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func string_to_int(s string) int {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}
	return int(i)
}

func main() {
	var start int = 1
	var sep int = 1
	var end int = 1

	delimeter := flag.String("s", "\n", "seperator of numbers")

	flag.Parse()
	args := flag.Args()
	var nopts int = len(args)

	switch nopts {
	case 1:
		end = string_to_int(args[0])
	case 2:
		start = string_to_int(args[0])
		end = string_to_int(args[1])
	case 3:
		start = string_to_int(args[0])
		sep = string_to_int(args[1])
		end = string_to_int(args[2])
	}

	for i := start; i <= end; i += sep {
		fmt.Printf("%d", i)
		if i+sep <= end {
			fmt.Printf("%s", *delimeter)
		} else {
			fmt.Printf("\n")
		}
	}
}
