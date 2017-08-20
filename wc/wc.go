package main

import (
	"bufio"
	"fmt"
	"github.com/pborman/getopt"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

type options struct {
	isBytes bool
	isChars bool
	isWords bool
	isLines bool
}

type counter struct {
	c int // bytes
	m int // characters
	l int // lines
	w int // words
}

func printCounts(o options, c counter) {
	if o.isLines {
		fmt.Printf("%d ", c.l)
	}
	if o.isWords {
		fmt.Printf("%d ", c.w)
	}
	if o.isChars {
		fmt.Printf("%d ", c.m)
	} else if o.isBytes {
		fmt.Printf("%d ", c.c)
	}
}

// words are define as any space delimited substring
func countWords(str string) int {
	word := regexp.MustCompile("\\S+")
	s := word.FindAllString(str, -1)
	return len(s)
}

func countChars(str string) int {
	return utf8.RuneCountInString(str)
}

// If no parameters are given default to showing lines, words and bytes
func setDefaultOptions(o *options) {
	if !o.isBytes && !o.isChars && !o.isWords && !o.isLines {
		o.isBytes = true
		o.isWords = true
		o.isLines = true
	}
}

func main() {
	var totalCount counter
	var stdinOnly = false

	c := getopt.Bool('c', "count bytes")
	m := getopt.Bool('m', "count chars")
	w := getopt.Bool('w', "count words")
	l := getopt.Bool('l', "count lines")
	getopt.Parse()

	opts := options{*c, *m, *w, *l}

	setDefaultOptions(&opts)

	args := getopt.Args()
	nargs := getopt.NArgs()

	/* Add empty file to list if its empty */
	if nargs == 0 {
		args = append(args, "")
		stdinOnly = true
	}

	/* Loop through the file reading the statistics */
	for _, file := range args {
		var f = os.Stdin
		if file != "-" && file != "" && !stdinOnly {
			var err error
			f, err = os.Open(file)
			if err != nil {
				os.Exit(1)
			}
		}

		var count counter
		var lastLine = false

		nr := bufio.NewReader(f)
		for {
			line, err := nr.ReadString('\n')
			if err == io.EOF {
				if len(line) == 0 {
					break
				}
				lastLine = true
			} else if err != nil {
				os.Exit(-1)
			}

			count.l++
			count.c += len(line)
			count.m += countChars(line)
			count.w += countWords(line)

			if lastLine {
				break
			}
		}

		/* Print the outcome */
		printCounts(opts, count)
		fmt.Println(file)

		/* Update total counts */
		if nargs > 1 {
			totalCount.c += count.c
			totalCount.m += count.m
			totalCount.w += count.w
			totalCount.l += count.l
		}
	}

	if nargs > 1 {
		/* Print the outcome */
		printCounts(opts, totalCount)
		fmt.Println("total")
	}
}
