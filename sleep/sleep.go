package main

import (
	"os"
	"time"
	"flag"
	"fmt"
	"strconv"
	"math"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: sleep time\n")
		os.Exit(1)
	}

	secs, err := strconv.Atof64(flag.Arg(0))

	if err != nil {
		fmt.Fprintf(os.Stderr, "sleep: Unknown Argument\n")
		os.Exit(1)
	}

	nsfp := secs * math.Pow10(9)

	nsi, _ := math.Modf(nsfp)

	time.Sleep(int64(nsi))
}
