package main

import (
	"os"
	"fmt"
	"strings"
	"flag"
)

var dflag = flag.Bool("d", false, "Print directory component")

func usage() {
	fmt.Fprintln(os.Stderr, "usage: basename [-d] string [suffix]")
	os.Exit(1)
}

func printDir(str string) {
	i := strings.LastIndex(str, "/")
	if i < 0 {
		fmt.Println(".")
		return
	}
	fmt.Println(str[:i])
}

func printName(str, suf string) {
	i, j := strings.LastIndex(str, "/"), strings.LastIndex(str, suf)
	fmt.Println(str[i+1:j])
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 || flag.NArg() > 2 {
		usage()
	}
	if *dirMode {
		printDir(flag.Arg(0))
	} else {
		if flag.NArg() == 2 {
			printName(flag.Arg(0), flag.Arg(1))
		} else {
			printName(flag.Arg(0), "")
		}
	}
}
