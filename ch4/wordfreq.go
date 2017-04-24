package ch4

import (
	"bufio"
	"fmt"
	"strings"
	// "os"
)

func Wordfreq() {
	count := make(map[string]int)
	const str = "test test done" // can be read from input file
	// scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(strings.NewReader(str))
	// data := input.Split(bufio.ScanWords)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// input := scanner.Text()
		count[scanner.Text()]++
		// fmt.Printf("input = %v\n", input)
	}
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	// }

	for k, v := range count {
		fmt.Printf("%s = %d\n", k, v)
	}
}
