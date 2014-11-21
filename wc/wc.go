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
func count_words(str string) int {
	word := regexp.MustCompile("\\S+")
	s := word.FindAllString(str, -1)
	return len(s)
}

func count_chars(str string) int {
	return utf8.RuneCountInString(str)
}

func main() {
	/* c = bytes
	 * m = chars
	 * l = lines
	 * w = words
	 */
	var total_c, total_m, total_l, total_w int = 0, 0, 0, 0
	var stdin_only bool = false

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

	/* Add emtpy file to list if its empty */
	if nargs == 0 {
		args = append(args, "")
		stdin_only = true
	}

	/* Loop through the file reading the statistics */
	for _, file := range args {
		var f *os.File = os.Stdin
		if file != "-" && file != "" && !stdin_only {
			var err error
			f, err = os.Open(file)
			if err != nil {
				os.Exit(1)
			}
		}

		var c, m, l, w int = 0, 0, 0, 0
		var last_line bool = false

		nr := bufio.NewReader(f)
		for {
			line, err := nr.ReadString('\n')
			if err == io.EOF {
				if len(line) == 0 {
					break
				}
				last_line = true
			} else if err != nil {
				os.Exit(-1)
			}
			l++
			if *isBytes {
				c += len(line)
				total_c += len(line)
			}
			if *isChars {
				m += count_chars(line)
			}
			if *isWords {
				w += count_words(line)
			}
			if last_line {
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
			total_c += c
			total_m += m
			total_w += w
			total_l += l
		}
	}

	if nargs > 1 {
		/* Print the outcome */
		if *isLines {
			fmt.Printf("%d ", total_l)
		}
		if *isWords {
			fmt.Printf("%d ", total_w)
		}
		if *isChars {
			fmt.Printf("%d ", total_m)
		} else if *isBytes {
			fmt.Printf("%d ", total_c)
		}
		fmt.Println("total")
	}
}
