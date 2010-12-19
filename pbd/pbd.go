package main

import (
	"os"
	"fmt"
)

func main() {
	var p, e = os.Getwd()

	if e != nil {
		fmt.Fprintf(os.Stderr, "pdb: can't find working directory: %s", e.String())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, p)
}
