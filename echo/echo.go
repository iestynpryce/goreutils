package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	newline := flag.Bool("n", false, "no trailing newline")

	flag.Parse()

	str := strings.Join(flag.Args(), " ")

	if !*newline {
		fmt.Println(str)
	} else {
		fmt.Print(str)
	}
}
