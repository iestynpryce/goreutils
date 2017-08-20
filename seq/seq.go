package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type testEnd func(int, int) bool

func moreThanEq(a int, b int) bool {
	return a >= b
}

func lessThanEq(a int, b int) bool {
	return a <= b
}

func stringToInt(s string) int {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}
	return int(i)
}

func main() {
	var start = 1
	var sep = 1
	var end = 1

	var t testEnd = lessThanEq

	delimeter := flag.String("s", "\n", "seperator of numbers")

	flag.Parse()
	args := flag.Args()
	var nopts = len(args)

	switch nopts {
	case 1:
		end = stringToInt(args[0])
	case 2:
		start = stringToInt(args[0])
		end = stringToInt(args[1])
	case 3:
		start = stringToInt(args[0])
		sep = stringToInt(args[1])
		end = stringToInt(args[2])
	}

	if start < end && sep < 0 {
		os.Exit(0)
	}
	if start > end && sep < 0 {
		t = moreThanEq
	}

	for i := start; t(i, end); i += sep {
		fmt.Printf("%d", i)
		if t(i+sep, end) {
			fmt.Printf("%s", *delimeter)
		} else {
			fmt.Printf("\n")
		}
	}
}
