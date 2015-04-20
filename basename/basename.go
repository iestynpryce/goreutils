// Implementation of basename
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var name string

func main() {
	name = os.Args[0]

	// Process command line flags
	sptr := flag.String("s", "", "suffix for removal")
	aptr := flag.Bool("a", false, "multiple arguments treated as NAME")
	flag.Parse()

	nargs := len(flag.Args())

	if *sptr != "" {
		if nargs < 1 {
			fmt.Fprintf(os.Stderr, "%s: missing operand\n", name)
			os.Exit(1)
		}
	}

	// Process multiple paths
	if *aptr || len(*sptr) > 0 {
		for _, arg := range flag.Args() {
			basename := filepath.Base(arg)
			basename = strings.TrimSuffix(basename, *sptr)
			fmt.Println(basename)
		}
	} else { // Process a single path
		if nargs > 2 {
			fmt.Fprintf(os.Stderr, "%s: extra operand: %v\n", name, flag.Args()[2:])
			os.Exit(1)
		}

		basename := filepath.Base(flag.Args()[0])
		if nargs == 2 {
			basename = strings.TrimSuffix(basename, flag.Args()[1])
		}

		fmt.Println(basename)
	}
}
