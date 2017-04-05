package ch2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit newline")
var sep = flag.String("s", " ", "separator")

func Echo4() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
