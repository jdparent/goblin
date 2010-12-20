package main

import (
	"os"
	"fmt"
	"time"
	"flag"
	"strconv"
)


func parseMonth(month string) int {
	switch month {
		case "jan", "january":
			return 1
		case "feb", "february":
			return 2
		case "mar", "march":
			return 3
		case "apr", "april":
			return 4
		case "may":
			return 5
		case "jun", "june":
			return 6
		case "jul", "july":
			return 7
		case "aug", "august":
			return 8
		case "sep", "september":
			return 9
		case "oct", "october":
			return 10
		case "nov", "november":
			return 11
		case "dec", "december":
			return 12
	}
	return 0
}

func parseYear(year string) int64 {
	var yr, er = strconv.Atoi(year)

	if er != nil {
		fmt.Fprintf(os.Stderr, "cal: error parsing parameters\n")
		os.Exit(1)
	}

	return int64(yr)
}

func januaryFirst(year int64) int {
	d := 4 + year + (year + 3) / 4

	if year > 1800 {
		d -= (year - 1701) / 100
		d += (year - 1601) / 400
	}

	if year > 1752 {
		d += 3
	}

	return int(d % 7)
}

var nonleapyear = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var leapyear = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func months(year int64) []int {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return leapyear
	}
	return nonleapyear
}

func printCal(month int, year int64) {
	var dayw string = " Su Mo Tu We Th Fr Sa"

	smon := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	var s string = smon[month-1] + " " + strconv.Itoa64(year)

	var off int = (20 - len(s)) / 2

	for i := 0; i < off; i++ {
		s = " " + s
	}

	s = s + "\n" + dayw + "\n"

	mth := months(year)

	day := januaryFirst(year)

	for i := 1; i < month; i++ {
		day += mth[i-1]
	}

	for i := 0; i < day % 7; i++ {
		s = s + "   "
	}

	for i := 1; i <= mth[month-1]; i++ {
		s = s + " " 
		if i < 10 {
			s = s + " "
		}
		s = s + strconv.Itoa(i)
		day += 1
		if day % 7 == 0 {
			s = s + "\n"
		}
	}

	s = s + "\n\n"

	os.Stdout.WriteString(s)
}

func main() {
	flag.Parse()

	var local = *time.LocalTime()

	var month int = 0
	var year int64 = 0

	if flag.NArg() > 2 {
		fmt.Fprintf(os.Stderr, "cal: error parsing parameters\n")
		os.Exit(1)
	}

	if flag.NArg() == 2 {
		month = parseMonth(flag.Arg(0))
		year = parseYear(flag.Arg(1))

		if month == 0 {
			fmt.Fprintf(os.Stderr, "cal: error parsing parameters\n")
			os.Exit(1)
		}
	} else {
		if flag.NArg() == 1 {
			year = local.Year
			month = parseMonth(flag.Arg(0))

			if month == 0 {
				year = parseYear(flag.Arg(0))
				month = local.Month
			}
		} else {
			year = local.Year
			month = local.Month
		}
	}

	printCal(month, year)
}
