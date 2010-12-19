// echo - print arguments
// 
// Code is pretty much line for line with the golang.org website's example app.
// Not a whole lot can be changed to get a more simple implementation.
package main

import (
	"os"
	"flag"
)

var suppressNewline = flag.Bool("n", false, "Suppress final newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if !*suppressNewline {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
