package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	pflag = flag.Bool("p", false, "Create parent directories")
	mflag = flag.String("m", "777", "Mode")
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: mkdir [ -p ] [ -m mode ] dirname ...")
	os.Exit(1)
}

func parseMode(perm string) (os.FileMode, error) {
	if len(perm) < 3 || len(perm) > 4 {
		fmt.Fprintf(os.Stderr, "mkdir: Unknown mode 1\n")
		os.Exit(1)
	}
	mode, err := strconv.ParseUint(perm, 8, 32)
	return os.FileMode(mode), err
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}
	mode, err := parseMode(*mflag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mkdir:  error parsing mode:", err.Error())
	}
	
	for _, name := range flag.Args() {
		if *pflag {
			if err := os.MkdirAll(name, mode); err != nil {
				fmt.Fprintf(os.Stderr, "mkdir: %s\n", err.Error())
			}
		} else {
			if err := os.Mkdir(name, mode); err != nil {
				fmt.Fprintf(os.Stderr, "mkdir: %s\n", err.Error())
			}
		}
	}
}
