package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	nLines := flag.Int("n", 10, "number of lines")
	flag.Parse()

	for _, file := range flag.Args() {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(-1)
		}

		var last_line bool = false
		var lines int = 0

		nr := bufio.NewReader(f)
		for {
			line, err := nr.ReadString('\n')
			if err == io.EOF {
				last_line = true
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
			lines++

			fmt.Print(line)

			if *nLines == lines || last_line {
				break
			}

		}
	}
}
