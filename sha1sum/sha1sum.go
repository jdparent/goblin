package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: sha1sum [file ...]\n")
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 { // Stdin
		f := os.Stdin
		h := sha1.New()
		var buf [64]byte

		for {
			switch nr, _ := f.Read(buf[:]); true {
			case nr > 0:
				_, e := h.Write(buf[0:nr])

				if e != nil {
					fmt.Fprintf(os.Stderr, "sha1sum: error generating hash\n")
					os.Exit(1)
				}
			case nr == 0:
				s := h.Sum(nil)

				for j := 0; j < len(s); j++ {
					if s[j] < 0x10 {
						fmt.Fprintf(os.Stdout, "0%0x", s[j])
					} else {
						fmt.Fprintf(os.Stdout, "%0x", s[j])
					}
				}
				fmt.Fprintf(os.Stdout, "  -\n")
				return
			case nr < 0:
				fmt.Fprintf(os.Stderr, "sha1sum: error generating hash\n")
				os.Exit(1)
			}
		}
	} else {
		h := sha1.New()

		for i := 0; i < flag.NArg(); i++ {
			d, e := ioutil.ReadFile(flag.Arg(i))

			if e != nil {
				fmt.Fprintf(os.Stderr, "sha1sum: cant read file %s\n", flag.Arg(i))
				os.Exit(1)
			}

			_, e = h.Write(d)

			if e != nil {
				fmt.Fprintf(os.Stderr, "sha1sum: error generating hash\n")
				os.Exit(1)
			}

			s := h.Sum(nil)

			for j := 0; j < len(s); j++ {
				if s[j] < 0x10 {
					fmt.Fprintf(os.Stdout, "0%0x", s[j])
				} else {
					fmt.Fprintf(os.Stdout, "%0x", s[j])
				}
			}

			fmt.Fprintf(os.Stdout, "  %s\n", flag.Arg(i))

			h.Reset() // Reset has for next file
		}
	}
}
