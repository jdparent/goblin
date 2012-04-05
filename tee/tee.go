package main

import (
	"container/vector"
	"flag"
	"fmt"
	"os"
)

var ignoreInterrupts = flag.Bool("i", false, "Ignore interrupts")
var appendOutput = flag.Bool("a", false, "Append the output to the files rather than rewriting them")

func main() {
	flag.Parse()

	var vec vector.Vector

	vec.Push(os.Stdout)

	if *ignoreInterrupts {
		fmt.Fprintf(os.Stdout, "tee: -i flag not supported\n")
		os.Exit(1)
	}

	for i := 0; i < flag.NArg(); i++ {
		if *appendOutput {
			f, e := os.OpenFile(flag.Arg(i), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

			if e != nil {
				fmt.Fprintf(
					os.Stderr,
					"tee: unable to open file %s: %s\n",
					flag.Arg(i),
					e.String())
				os.Exit(1)
			}

			vec.Push(f)
		} else {
			f, e := os.Create(flag.Arg(i))

			if e != nil {
				fmt.Fprintf(
					os.Stderr,
					"tee: unable to open file %s: %s\n",
					flag.Arg(i),
					e.String())
				os.Exit(1)
			}

			vec.Push(f)
		}
	}

	// Copied from cat.go
	var buf [64]byte
	f := os.Stdin

	for {
		switch nr, _ := f.Read(buf[:]); true {
		case nr > 0:
			for x := 0; x < len(vec); x++ {
				el := vec.At(x).(*os.File)
				_, e := el.Write(buf[0:nr])

				if e != nil {
					fmt.Fprintf(
						os.Stderr,
						"tee: error writting to file %s: %s\n",
						el.Name(),
						e.String())
					os.Exit(1)
				}
			}
		case nr == 0:
			for x := 0; x < len(vec); x++ {
				el := vec.At(x).(*os.File)
				el.Close()
			}
			return
		case nr < 0:
			fmt.Fprintf(os.Stderr, "tee: error reading from stdin\n")
			os.Exit(1)
		}
	}
}
