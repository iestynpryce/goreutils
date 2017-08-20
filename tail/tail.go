package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

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

		if !*quiet && numfiles > 1 {
			fmt.Printf("==> %s <==\n", file)
		}

		/* Read the entire file into memory, then split on newlines into an array */
		fileAll, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(-1)
		}
		fileArray := strings.SplitAfter(string(fileAll), "\n")

		var fileLen = len(fileArray)

		/* Print the last N lines */
		for i = max(fileLen-*nLines-1, 0); i < fileLen; i++ {
			line := fileArray[i]
			fmt.Print(line)
		}

		/* Print multi file seperator */
		if !*quiet && numfiles > 1 && i < len(flag.Args())-1 {
			fmt.Println("")
		}
	}
}
