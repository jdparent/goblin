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
				fmt.Fprintln(os.Stderr, "cat:", werr.Error())
				os.Exit(1)
			}
		case nr == 0:
			return
		case nr < 0:
			fmt.Fprintln(os.Stderr, "cat:", rerr.Error())
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
			if err != nil {
				fmt.Fprintln(os.Stderr, "cat:" err.Error())
				os.Exit(1)
			}
			cat(f)
			f.Close()
		}
	}
}
