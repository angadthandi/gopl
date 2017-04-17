package ch4

import (
	"fmt"
	"unicode"
)

func Reverse() {
	s := []int{1, 2, 3, 4, 5}                // slice
	a := [...]int{1, 2, 3, 4, 5}             // array
	d := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5} // slice

	str := "1  2  3"
	reverse(s)
	reversePtr(&a)
	d = replaceDups(d)
	str = squashSpace(str)
	// fmt.Printf("%v\n%T\n", s, s)
	// fmt.Printf("%v\n%T\n", a, a)
	// fmt.Printf("%v\n%T\n", d, d)
	fmt.Printf("%v\n%T\n", str, str)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reversePtr(a *[5]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func replaceDups(s []int) []int {
	r, tmp := s[:0], 0
	for i := 0; i < len(s); i = i + 1 {
		if s[i] != tmp {
			r = append(r, s[i])
		}
		tmp = s[i]
	}

	return r
}

func squashSpace(str string) string {
	s := []rune(str)
	c := 0
	for i := 0; i < len(s); i = i + 1 {
		if unicode.IsSpace(s[i]) {
			c++
			if c > 1 {
				s[i] = rune(0)
				c--
			}
		} else if c > 0 {
			c--
		}
	}

	return string(s)
}
