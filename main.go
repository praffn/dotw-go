package main

import (
	"flag"
	"fmt"
	"github.com/praffn/dotw/dotw"
	"time"
)

func main() {
	year := time.Now().Year()
	fromYear := flag.Int(
		"from",
		year,
		fmt.Sprintf("the from year, default is %d", year),
	)
	toYear := flag.Int(
		"to",
		year,
		fmt.Sprintf("the to year, default is %d", year),
	)
	flag.Parse()

	if *fromYear > *toYear {
		panic("lol m8")
	}

	for {
		dotw.Play(*fromYear, *toYear)
	}
}
