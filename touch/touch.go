// touch - set modification date of a file
package main

import (
	"os"
	"flag"
	"fmt"
	"syscall"
	"time"
)

var create = flag.Bool("c", false, "Don't create if not exists")
var newTime = flag.Int("t", int(time.Seconds()), "Set to time provided")

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

	for i := 0; i < flag.NArg(); i++ {
		var name = flag.Arg(i)
		var tb syscall.Utimbuf
		tb.Actime = (int32)(*newTime)
		tb.Modtime = (int32)(*newTime)
		var e = syscall.Utime(name, &tb)

		if (e != 0) && *create {
			fmt.Fprintf(
				os.Stderr,
				"touch: cannot touch `%s'\n",
				name)
			os.Exit(1)
		}

		f, err := os.Open(name, os.O_CREAT, 0666)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"touch: cannot touch `%s': %s\n",
				name,
				err.String())
			os.Exit(1)
		}
		f.Close()

		e = syscall.Utime(name, &tb)
		if e != 0{
			fmt.Fprintf(
			os.Stderr,
			"touch: cannot touch `%s'\n",
			name)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
