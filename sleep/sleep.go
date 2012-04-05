package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: sleep time\n")
		os.Exit(1)
	}

	secs, err := strconv.ParseFloat(flag.Arg(0), 64)

	if err != nil {
		fmt.Fprintf(os.Stderr, "sleep: Unknown Argument\n")
		os.Exit(1)
	}

	nsfp := secs * math.Pow10(9)

	nsi, _ := math.Modf(nsfp)

	time.Sleep(int64(nsi))
}
