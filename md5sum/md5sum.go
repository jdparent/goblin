package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: md5sum [file ...]\n")
	os.Exit(1)
}

func main() {
	flag.Parse()
	var buf [64]byte
	h := md5.New()
	
	if flag.NArg() == 0 {
		for {
			nr, rerr := os.Stdin.Read(buf[:])
			switch {
			case nr > 0:
				// h.Write never returns an error.
				h.Write(buf[0:nr])
			case nr == 0:
				fmt.Printf("%x\n", h.Sum(nil))
				os.Exit(0)
			case nr < 0:
				fmt.Fprintf(os.Stderr, "md5sum: error reading from os.Stdin: %s\n", rerr)
				os.Exit(1)
			}
		}
	} else {
		for _, v := range flag.Args() {
			data, err := ioutil.ReadFile(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "md5sum: error reading file %s: %s\n", v, err)
				os.Exit(1)
			}

			// h.Write never returns an error.
			h.Write(data)
			fmt.Printf("%x %s\n", h.Sum(nil), v)
			h.Reset()
		}
	}
}
