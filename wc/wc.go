package main

import (
	"bufio"
	"code.google.com/p/getopt"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
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
	var c, m, l, w, L int = 0, 0, 0, 0, 0

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

	/* Loop through the file reading the statistics */
	for _, file := range getopt.Args() {
		f, err := os.Open(file)
		if err != nil {
			os.Exit(1)
		}

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
				os.Exit(1)
			}
			l++
			tmp_l := len(strings.TrimSuffix(line, "\n"))
			if tmp_l > L {
				L = tmp_l
			}
			if *isBytes {
				c += len(line)
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
	}
}
