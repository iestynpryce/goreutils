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

	var numfiles int = len(flag.Args())

	for i, file := range flag.Args() {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(-1)
		}

		var last_line bool = false
		var lines int = 0

		if !*quiet && numfiles > 1 {
			fmt.Printf("==> %s <==\n", file)
		}

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

		/* Print multi file seperator */
		if !*quiet && numfiles > 1 && i < len(flag.Args())-1 {
			fmt.Println("")
		}
	}
}
