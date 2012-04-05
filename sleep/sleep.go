package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: sleep time\n")
		os.Exit(1)
	}

	if t, err := time.ParseDuration(flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "sleep: Unknown Argument\n")
		os.Exit(1)
	} else {
		time.Sleep(t)
	}
}
