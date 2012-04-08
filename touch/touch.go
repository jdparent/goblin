// touch - set modification date of a file
package main

import (
	"os"
	"flag"
	"fmt"
	"time"
)

var (
	aflag = flag.Bool("a", false, "Change access time instead of modification time")
	tflag = flag.Int("t", 0, "Set to time provided")
)

func usage() {
	fmt.Fprintf(
		os.Stderr,
		"usage: touch [-c] [-t time] names...\n")
	os.Exit(1)
}

func main() {
	flag.Parse();
	if flag.NArg() < 1 {
		usage()
	}
	var atime, mtime time.Time
	if *tflag == 0 {
		if *aflag {
			atime = time.Now()
		} else {
			mtime = time.Now()
		}
	} else {
		if *aflag {
			atime = time.Unix(tflag, 0)
		} else {
			mtime = time.Unix(tflag, 0)
		}
	}
	for _, name := range flag.Args() {
		if err := os.Chtimes(name, atime, mtime); err != nil {
			fmt.Fprintln(os.Stderr, "touch: cannon touch %s: %s", name, err.Error())
		}
	}
}