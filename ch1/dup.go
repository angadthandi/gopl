package ch1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Dup1() {
	fmt.Println("Dup1...")
	start1 := time.Now()
	// fmt.Println(os.Stdin)
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// fmt.Println(input)
	for input.Scan() {
		// fmt.Println(input.Text())
		counts[input.Text()]++
	}
	// fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			// fmt.Println(n)
			// fmt.Println(line)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	secs1 := time.Since(start1).Seconds()
	fmt.Println(secs1)
}

func Dup2() {
	fmt.Println("Dup2...")
	start2 := time.Now()

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(counts) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, n := range counts {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	secs2 := time.Since(start2).Seconds()
	fmt.Println(secs2)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func Dup3() {
	fmt.Println("Dup3...")
	start3 := time.Now()

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	secs3 := time.Since(start3).Seconds()
	fmt.Println(secs3)
}
