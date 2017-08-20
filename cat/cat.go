package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Global for program name
var name string

// Error checking helper
func checkAndExit(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, e)
		os.Exit(1)
	}
}

// Print lines according to the options given
func printLine(n bool, str string, line int) {
	if n {
		fmt.Printf("%6d  ", line)
	}
	fmt.Print(str)
}

func main() {
	name = os.Args[0]

	// Process cmd line flags
	nptr := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	// Initialise line count
	var linecount int

	// String buffer to store line contents for strings
	// not ending in a new line
	var strBuf string

	for _, file := range flag.Args() {
		f, err := os.Open(file)
		checkAndExit(err)
		reader := bufio.NewReader(f)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				strBuf = str
				break
			} else {
				checkAndExit(err)
			}

			linecount++
			printLine(*nptr, strBuf+str, linecount)

			// Reset the string buffer as it's now been printed
			strBuf = ""
		}
	}

	// Write out the string buffer if it contains anything
	if len(strBuf) > 0 {
		linecount++
		printLine(*nptr, strBuf, linecount)
	}
}
