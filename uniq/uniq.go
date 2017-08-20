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

// Open the input and output file descriptors
func openFiles(args []string) (*bufio.Reader, *os.File) {
	var reader *bufio.Reader
	var out = os.Stdout

	if len(args) == 0 {
		return bufio.NewReader(os.Stdin), out
	}

	if len(args) > 1 {
		f, err := os.OpenFile(args[1],
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
			os.FileMode(0666))
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		out = f
	}

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	reader = bufio.NewReader(f)

	return reader, out
}

func main() {
	var reader *bufio.Reader
	var out *os.File
	var firstLine = true
	var lastLine string

	c = flag.Bool("c", false, "precede each output line with a count of the number of times it occurred")
	d := flag.Bool("d", false, "only print duplicate lines, one for each group")
	u := flag.Bool("u", false, "only print non duplicate lines")

	flag.Parse()
	args := flag.Args()

	// Open stdin or provided file as input, and stdout or provided file as output
	reader, out = openFiles(args)

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
			} else if !(*d && *u) { // one of d and u are true, but not both
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

	if err := out.Close(); err != nil {
		log.Fatalf("Error %v\n", err)
	}

	os.Exit(0)
}
