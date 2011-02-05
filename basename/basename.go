package main

import (
	"os"
	"fmt"
	"strings"
	"flag"
)

var dirMode = flag.Bool("d", false, "Print directory component")

func usage() {
	fmt.Fprintf(
		os.Stderr,
		"usage: basename [-d] string [suffix]\n")
	os.Exit(1)
}

func printDir(str string) {
	index := strings.LastIndex(str, "/")

	if index == -1 {
		fmt.Fprintf(os.Stdout, ".\n")
	} else {
		if index == 0 {
			fmt.Fprintf(os.Stdout, "\n")
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", str[0:index])
		}
	}
}

func printName(str, suf string) {
	index := strings.LastIndex(str, "/")
	name := str[index+1:len(str)]

	if len(suf) > 0 {
		index2 := strings.LastIndex(name, suf)

		if index2 != -1 {
			name = name[0:index2]
		}
	}
	fmt.Fprintf(os.Stdout, "%s\n", name)
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
