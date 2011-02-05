package main

import (
	"os"
	"fmt"
	"strings"
	"flag"
	"path"
)

var pwd = flag.String("d", "", "Sets root for unrooted names")

func usage() {
	fmt.Fprintf(
		os.Stderr,
		"usage: cleanname [-d pwd] names...\n")
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
	}

	for i := 0; i < flag.NArg(); i++ {
		var name = flag.Arg(i)
		if strings.HasPrefix(name, "/") == false {
			if *pwd != "" {
				name = *pwd + "/" +  name
			}
		}
		name = path.Clean(name)
		fmt.Fprintf(os.Stdout, "%s\n", name)
	}
}
