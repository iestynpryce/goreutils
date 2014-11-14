package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Error checking helper
func check_and_exit(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", e)
		os.Exit(-1)
	}
}

// Print the file in the reverse order using recursive functions
func reverse_print(r *bufio.Reader) {
	str, err := r.ReadString('\n')
	if err == io.EOF {
		fmt.Print(str)
		return
	} else {
		check_and_exit(err)
	}
	reverse_print(r)
	fmt.Print(str)
}

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		f, err := os.Open(file)
		check_and_exit(err)
		reader := bufio.NewReader(f)
		reverse_print(reader)
	}
}
