package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	i := flag.Int("i", 1, "line-increment")
	s := flag.String("s", "\t", "number seperator")

	flag.Parse()

	args := flag.Args()

	var line_num int = 0
	for _, file := range args {

		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		r := bufio.NewReader(f)
		for {
			var last_line bool = false
			line, err := r.ReadString('\n')
			if err == io.EOF {
				if len(line) == 0 {
					break
				}
				last_line = true
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				break
			}

			// Only iterate and print line number if it's a non-blank line
			if len(strings.TrimSpace(line)) > 0 {
				line_num += *i
				fmt.Printf("%6d%s", line_num, *s)
			}

			fmt.Printf("%s", line)

			if last_line {
				break
			}
		}
	}
}
