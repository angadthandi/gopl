package ch1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Echo1() {
	start1 := time.Now()

	fmt.Println("echo1...")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("for : " + s)

	secs1 := time.Since(start1).Seconds()
	fmt.Println(secs1)
}

func Echo2() {
	start2 := time.Now()

	fmt.Println("echo2...")
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println("range : " + s2)

	secs2 := time.Since(start2).Seconds()
	fmt.Println(secs2)
}

func Echo3() {
	start3 := time.Now()

	fmt.Println("echo3...")
	fmt.Println("join : " + strings.Join(os.Args[1:], " "))

	secs3 := time.Since(start3).Seconds()
	fmt.Println(secs3)
}
