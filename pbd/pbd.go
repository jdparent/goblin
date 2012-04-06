package main

import (
	"fmt"
	"os"
)

func main() {

	if p, err := os.Getwd(); err != nil {
		fmt.Fprintln(os.Stderr, "pdb: can't find working directory:", err.Error())
		os.Exit(1)
	} else {
		fmt.Print(p)
	}
}
