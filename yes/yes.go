package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string

	if len(os.Args) > 1 {
		str = strings.Join(os.Args[1:], " ")
	} else {
		str = "y"
	}

	for {
		fmt.Println(str)
	}
}
