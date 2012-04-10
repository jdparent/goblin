// tee - pipe fitting

package main

import (
	"flag"
	"fmt"
	"os"
)

var aflag = flag.Bool("a", false, "Append the output to the files rather than rewriting them")

func main() {
	flag.Parse()

	files := make([]*os.File, 0, flag.NArg() + 1)
	files = append(files, os.Stdout)

	for _, v := range flag.Args() {
		var f *os.File
		var err error
		
		if *aflag {
			f, err = os.OpenFile(v, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		} else {
			f, err = os.Create(v)
		}
		defer f.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "tee: unable to open file %s: %s\n", v, err.Error())
			os.Exit(1)
		}
		files = append(files, f)
	}

	var buf [64]byte
	f := os.Stdin

	for {
		switch nr, rerr := f.Read(buf[:]); true {
		case nr > 0:
			for _, f := range files {
				_, err := f.Write(buf[0:nr])

				if err != nil {
					fmt.Fprintf(os.Stderr, "tee: error writting to file %s: %s\n", f.Name(), err.Error())
					os.Exit(1)
				}
			}
		case nr == 0:
			os.Exit(0)
		case nr < 0:
			fmt.Fprintln(os.Stderr, "tee: error reading from stdin:", rerr)
			os.Exit(1)
		}
	}
}
