package ch2

import (
	"fmt"
	"github.com/angadthandi/gopl/ch2/generalconv"
)

func Tempconv() {
	fmt.Printf("vvhhh! %v\n", generalconv.AbsoluteZeroC)
	fmt.Println(generalconv.CToF(generalconv.BoilingC))
	fmt.Printf("CToK = %v\n", generalconv.CToK(generalconv.BoilingC))
}
