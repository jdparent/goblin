package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	uflag = flag.Bool("u", false, "Report Coordinated Universal Time (UTC) rather than local time.")
	nflag = flag.Bool("n", false, "Report the date as the number of seconds since the epoch, 00:00:00 UTC, January 1, 1970")
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: date [ -u ] [ -n ] [ seconds ]")
	os.Exit(1)
}

func main() {
	flag.Parse()
	var t time.Time

	switch flag.NArg() {
	case 0:
		t = time.Now()
	case 1:
		i, err := strconv.ParseInt(flag.Arg(0), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "date: error parsing time:", err.Error())
			os.Exit(1)
		}
		t = time.Unix(i, 0)
	default:
		usage()
	}

	switch {
	case *nflag:
		fmt.Println(t.Unix())
	case *uflag:
		fmt.Println(t.UTC())
	default:
		fmt.Println(t)
	}
}
