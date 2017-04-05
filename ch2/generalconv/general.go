package generalconv

import "fmt"

type Celsius float64
type Farenheit float64
type Kelvin float64

type Feet float64
type Metre float64

type Kilogram float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

func FToc(f Farenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(k Celsius) Kelvin {
	return Kelvin(k - (-273.15))
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func FeetToMetre(f Feet) Metre {
	return Metre(f / (5 / 2))
}

func KiloToPound(k Kilogram) Pound {
	return Pound(k / (5 / 2))
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Metre) String() string {
	return fmt.Sprintf("%gm", m)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%goz", p)
}
