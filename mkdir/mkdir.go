package main

import (
	"os"
	"fmt"
	"flag"
	"strconv"
)

var makeParents = flag.Bool("p", false, "Create parent directories")
var mode = flag.String("m", "777", "Mode")

func parseMode(perm string) uint32 {

	if len(perm) < 3 || len(perm) > 4 {
		fmt.Fprintf(
			os.Stderr,
			"mkdir: Unknown mode 1\n")
		os.Exit(1)
	}

	if len(perm) == 3 {
		perm = "0" + perm
	}

	newMode, cerr := strconv.Btoui64(perm, 8)

	if cerr != nil {
		fmt.Fprintf(
			os.Stderr,
			"mkdir: Unknown mode\n")
		os.Exit(1)
	}

	return uint32(newMode)
}

func main() {
	flag.Parse()

	var nm uint32 = parseMode(*mode)

	if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			if *makeParents {
				var err = os.MkdirAll(flag.Arg(i), 0777)
				if err != nil {
					fmt.Fprintf(
						os.Stderr,
						"mkdir: %s\n",
						err.String())
				}
			} else {
				var err = os.Mkdir(flag.Arg(i), nm)
				if err != nil {
					fmt.Fprintf(
						os.Stderr,
						"mkdir: %s\n",
						err.String())
				}
			}

			os.Chmod(flag.Arg(i), nm)
		}
	}
}
