package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Error checking helper
func checkAndExit(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", e)
		os.Exit(-1)
	}
}

// Print the file in the reverse order using recursive functions
func reversePrint(r *bufio.Reader) {
	str, err := r.ReadString('\n')
	if err == io.EOF {
		fmt.Print(str)
		return
	}

	checkAndExit(err)

	reversePrint(r)
	fmt.Print(str)
}

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		f, err := os.Open(file)
		checkAndExit(err)
		reader := bufio.NewReader(f)
		reversePrint(reader)
	}
}
