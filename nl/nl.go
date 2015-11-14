// nl: place line numbers at the beginning of lines
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	i := flag.Int("i", 1, "line-increment")
	s := flag.String("s", "\t", "number seperator")

	flag.Parse()

	args := flag.Args()

	if *i <= 0 {
		log.Fatalf("Error: invalid line-increment value \"%d\"\n", *i)
	}

	var lineNum = 1
	for _, file := range args {

		f, err := os.Open(file)
		if err != nil {
			log.Fatalln("Error: %v\n", err)
		}

		r := bufio.NewReader(f)
		for {
			var lastLine = false
			line, err := r.ReadString('\n')
			if err == io.EOF {
				if len(line) == 0 {
					break
				}
				lastLine = true
			} else if err != nil {
				log.Fatalln("Error: %v\n", err)
			}

			isBlank := len(strings.TrimSpace(line)) == 0

			// Only print line number if it's a non-blank line
			if !isBlank {
				fmt.Printf("%6d%s", lineNum, *s)
			}

			fmt.Printf("%s", line)

			// Only iterate line number if it's a non-blank line
			if !isBlank {
				lineNum += *i
			}

			if lastLine {
				break
			}
		}
	}
}
