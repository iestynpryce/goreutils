package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

// Global for program name
var name string

var wrapColumns int = 76

// Error checking helper
func checkAndExit(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, e)
		os.Exit(1)
	}
}

// Wrapped print
func wrapPrint(s string, w int) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
		if (i+1)%w == 0 {
			fmt.Print("\n")
		}
	}
}

func main() {
	name = os.Args[0]

	var buffer bytes.Buffer
	var str string

	// Process cmd line flags
	dptr := flag.Bool("d", false, "Decode")
	flag.Parse()

	for _, file := range flag.Args() {
		f, err := os.Open(file)
		checkAndExit(err)
		reader := bufio.NewReader(f)
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)

			if err == io.EOF {
				break
			} else {
				checkAndExit(err)
			}
			buffer.Write(buf[:n])
		}
	}
	if *dptr {
		// Decode
		data, err := base64.StdEncoding.DecodeString(buffer.String())
		checkAndExit(err)
		fmt.Printf("%s", data[:])
	} else {
		// Encode
		str = base64.StdEncoding.EncodeToString(buffer.Bytes())
		wrapPrint(str, wrapColumns)
		fmt.Print("\n")
	}
}
