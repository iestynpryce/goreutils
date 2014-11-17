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

	var line_num int = 1
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

			is_blank := len(strings.TrimSpace(line)) > 0

			// Only print line number if it's a non-blank line
			if is_blank {
				fmt.Printf("%6d%s", line_num, *s)
			}

			fmt.Printf("%s", line)

			// Only iterate line number if it's a non-blank line
			if is_blank {
				line_num += *i
			}

			if last_line {
				break
			}
		}
	}
}
