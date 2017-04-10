package ch3

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func Comma() {
	for _, v := range os.Args[1:] {
		c := comma(v)
		fmt.Printf("v = %s\n", c)
	}
}

func comma(s string) string {
	sep, sfx := ".", ""
	if strings.Contains(s, sep) {
		sS := strings.Split(s, sep)

		if len(sS[0]) <= 3 {
			return s
		}
		s = sS[0]
		sfx = sep + sS[1]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	// return comma(s[:n-3]) + "," + s[n-3:]

	var b bytes.Buffer
	b.Write([]byte(s[:n-3]))
	b.Write([]byte(","))
	b.Write([]byte(s[n-3:]))
	b.Write([]byte(sfx))

	return b.String()
}

func Anagram() {
	s1 := os.Args[1]
	s2 := os.Args[2]

	r := anagram(s1, s2)

	fmt.Printf("r = %g\n", r)
}

func anagram(s1, s2 string) bool {
	c := 0

	n1 := len(s1)
	n2 := len(s2)

	if n1 != n2 {
		return false
	}
	b1 := []byte(s1)
	b2 := []byte(s2)
	for _, v1 := range b1 {
		for _, v2 := range b2 {
			if v1 == v2 {
				c++
			}
		}
	}

	if c == n1 {
		return true
	}

	return false
}
