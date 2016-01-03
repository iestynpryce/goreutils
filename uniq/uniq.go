// uniq: remove repeated adjacent lines
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var reader *bufio.Reader
	var out *os.File
	var firstLine bool = true
	var lastLine string

	d := flag.Bool("d", false, "only print duplicate lines, one for each group")

	flag.Parse()
	args := flag.Args()

	// Open stdin or provided file as input
	if len(args) == 0 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		reader = bufio.NewReader(f)

		if len(args) > 1 {
			f, err := os.OpenFile(args[1],
				os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
				os.FileMode(0666))
			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}
			out = f
			defer out.Close()
		}
	}

	// Default to outputting to stdout
	if out == nil {
		out = os.Stdout
	}

	// Loop over provided input
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			if len(line) == 0 {
				break
			}
		} else if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		if firstLine {
			lastLine = line
			firstLine = false
			if !*d {
				fmt.Fprintf(out, "%s", line)
			}
			continue
		}

		if line != lastLine {
			if !*d {
				fmt.Fprintf(out, "%s", line)
			}
		} else if *d {
			fmt.Fprintf(out, "%s", line)
		}

		lastLine = line
	}
}
