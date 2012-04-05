// cat - catenate files
//
// Code is pretty much line for line with the golang.org website's example 
// rotting cats.  Not a while lot can be changed to get a more simple 
// implementation.
package main

import (
	"flag"
	"fmt"
	"os"
)

const NBUF = 8192

func cat(f *os.File) {
	var b[NBUF]byte

	for {
		nr, rerr := f.Read(b[:])
		switch {
		case nr > 0:
			nw, werr := os.Stdout.Write(b[0:nr])
			if nw != nr {
				fmt.Fprintf(os.Stderr, "cat: write error copying %s: %s", f.Name(), werr.Error())
				os.Exit(1)
			}
		case nr == 0:
			return
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading %s: %s", f.Name(), rerr.Error())
			os.Exit(1)
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)		
	} else {
		for _, v := range flag.Args() {
			f, err := os.Open(v)
			if f == nil {
				fmt.Fprintf(os.Stderr, "cat: can't open %s: %s", v, err.Error())
				os.Exit(1)
			}
			cat(f)
			f.Close()
		}
	}
}
