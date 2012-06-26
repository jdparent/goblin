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

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		io.copy(os.Stdin, os.Stdout)		
	} else {
		for _, v := range flag.Args() {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintln(os.Stderr, "cat:" err.Error())
				os.Exit(1)
			}
			io.copy(f, os.Stdout)
			f.Close()
		}
	}
}
