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
	var firstLine bool = true
	var lastLine string

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		reader = bufio.NewReader(f)
	}

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
			fmt.Fprintf(os.Stdout, "%s", line)

			continue
		}

		if line != lastLine {
			fmt.Fprintf(os.Stdout, "%s", line)
		}

		lastLine = line
	}
}
