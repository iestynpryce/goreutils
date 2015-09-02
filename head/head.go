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
	quiet := flag.Bool("q", false, "supress file headers")
	flag.Parse()

	var numfiles = len(flag.Args())

	for i, file := range flag.Args() {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(-1)
		}
		defer f.Close()

		var lastLine = false
		var lines

		if !*quiet && numfiles > 1 {
			fmt.Printf("==> %s <==\n", file)
		}

		nr := bufio.NewReader(f)
		for {
			line, err := nr.ReadString('\n')
			if err == io.EOF {
				lastLine = true
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}

			lines++

			fmt.Print(line)

			if *nLines == lines || lastLine {
				break
			}

		}

		/* Print multi file seperator */
		if !*quiet && numfiles > 1 && i < len(flag.Args())-1 {
			fmt.Println("")
		}
	}
}
