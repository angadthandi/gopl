package ch2

import (
	"fmt"
	"os"
	"strconv"

	"github.com/angadthandi/gopl/ch2/generalconv"
)

func Cf() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf : %v\n", err)
		}
		c := generalconv.Celsius(t)
		f := generalconv.Farenheit(t)

		fmt.Printf("%s = %s, %s = %s",
			f, generalconv.FToc(f), c, generalconv.CToF(c))
	}
}

func GeneralConv() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "General : %v\n", err)
		}
		f := generalconv.Farenheit(t)
		c := generalconv.Celsius(t)

		l := generalconv.Feet(t)
		w := generalconv.Kilogram(t)

		fmt.Printf("%s = %s, %s = %s,\n %s = %s, %s = %s\n",
			f, generalconv.FToc(f), c, generalconv.CToF(c),
			l, generalconv.FeetToMetre(l), w, generalconv.KiloToPound(w))
	}
}
