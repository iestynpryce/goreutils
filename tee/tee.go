// tee: output stdin to file and stdout
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
	a := flag.Bool("a", false, "append output")

	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		reader := bufio.NewReader(os.Stdin)

		mode := os.FileMode(0666)
		permission := os.O_WRONLY | os.O_CREATE

		if *a {
			permission |= os.O_APPEND
		}

		f, err := os.OpenFile(args[0], permission, mode)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		defer f.Close()

		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				if len(line) == 0 {
					break
				}
			} else if err != nil {
				log.Fatalf("Error: %v\n", err)
			}

			fmt.Fprintf(os.Stdout, "%s", line)
			fmt.Fprintf(f, "%s", line)
		}
	}
}
