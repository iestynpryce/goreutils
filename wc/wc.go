package main

import (
	"bufio"
	"code.google.com/p/getopt"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

// words are define as any space delimited substring
func countWords(str string) int {
	word := regexp.MustCompile("\\S+")
	s := word.FindAllString(str, -1)
	return len(s)
}

func countChars(str string) int {
	return utf8.RuneCountInString(str)
}

func main() {
	/* c = bytes
	 * m = chars
	 * l = lines
	 * w = words
	 */
	var totalC, totalM, totalL, totalW int = 0, 0, 0, 0
	var stdinOnly = false

	isBytes := getopt.Bool('c', "count bytes")
	isChars := getopt.Bool('m', "count chars")
	isWords := getopt.Bool('w', "count words")
	isLines := getopt.Bool('l', "count lines")

	getopt.Parse()

	// If no parameters are given default to showing lines, words and bytes
	if !*isBytes && !*isChars && !*isWords && !*isLines {
		*isBytes = true
		*isWords = true
		*isLines = true
	}

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

		var c, m, l, w int = 0, 0, 0, 0
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
			l++
			if *isBytes {
				c += len(line)
				totalC += len(line)
			}
			if *isChars {
				m += countChars(line)
			}
			if *isWords {
				w += countWords(line)
			}
			if lastLine {
				break
			}
		}

		/* Print the outcome */
		if *isLines {
			fmt.Printf("%d ", l)
		}
		if *isWords {
			fmt.Printf("%d ", w)
		}
		if *isChars {
			fmt.Printf("%d ", m)
		} else if *isBytes {
			fmt.Printf("%d ", c)
		}
		fmt.Println(file)

		/* Update total counts */
		if nargs > 1 {
			totalC += c
			totalM += m
			totalW += w
			totalL += l
		}
	}

	if nargs > 1 {
		/* Print the outcome */
		if *isLines {
			fmt.Printf("%d ", totalL)
		}
		if *isWords {
			fmt.Printf("%d ", totalW)
		}
		if *isChars {
			fmt.Printf("%d ", totalM)
		} else if *isBytes {
			fmt.Printf("%d ", totalC)
		}
		fmt.Println("total")
	}
}
