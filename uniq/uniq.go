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

var (
	counter int = 1 // default to 1 as this is the minimum value
	c       *bool
)

func printLine(out io.Writer, line string) {
	if *c {
		fmt.Fprintf(out, "%d %s", counter, line)
	} else {
		fmt.Fprintf(out, "%s", line)
	}
}

func main() {
	var reader *bufio.Reader
	var out *os.File
	var firstLine bool = true
	var lastLine string

	c = flag.Bool("c", false, "precede each output line with a count of th e number of times it occurred")
	d := flag.Bool("d", false, "only print duplicate lines, one for each group")
	u := flag.Bool("u", false, "only print non duplicate lines")

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
				if !(*d && *u) {
					if !(*u && counter > 1) && !(*d && counter == 1) {
						printLine(out, lastLine)
					}
				}
				break
			}
		} else if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		if firstLine {
			lastLine = line
			firstLine = false
			continue
		}

		if line != lastLine {
			if !*d && !*u {
				printLine(out, lastLine)
			} else if !(*d && *u) {

				if *d && counter > 1 {
					printLine(out, lastLine)
				} else if *u && counter == 1 {
					printLine(out, lastLine)
				}
			}
			counter = 1 // reset to the minimum value
		} else {
			counter++
		}

		lastLine = line
	}

	os.Exit(0)
}
