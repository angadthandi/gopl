package ch4

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func ShaDiffBits() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	var ct1, ct2, rc int
	for _, v1 := range c1 {
		ct1++
		for _, v2 := range c2 {
			ct2++
			if v1 != v2 {
				rc++
			}
		}
	}

	fmt.Printf("Count1 = %v\n", ct1)
	fmt.Printf("Count2 = %v\n", ct2)
	fmt.Printf("Different Bytes Count = %v\n", rc)
}

func ShaOut() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		if args[1] == "SHA384" {
			r := sha512.Sum384([]byte(args[0]))
			fmt.Printf("%x\n", r)
		} else if args[1] == "SHA512" {
			r := sha512.Sum512([]byte(args[0]))
			fmt.Printf("%x\n", r)
		}
	} else {
		r := sha512.Sum512_256([]byte(args[0]))
		fmt.Printf("%x\n", r)
	}
}
