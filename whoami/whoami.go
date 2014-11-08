package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	usrptr, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get current user: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(usrptr.Username)
}
