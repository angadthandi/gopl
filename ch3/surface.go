package ch3

import (
	// "encoding/xml"
	"fmt"
	// "io"
	"math"
)

const (
	width, height = 600, 320 // canvas size px
	cells         = 100      // num of grids
	xyrange       = 30.0     // axis range
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func Surface() string {
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// need to revise coloring -incorrect
			clr := "#0000ff" // blue for valleys
			if f(ax, ay) > f(cx, cy) {
				clr = "#ff0000" // red for peaks
			}

			svg += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+
				"style='stroke: %s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, clr)
		}
	}
	svg += fmt.Sprintf("%s", "</svg>")

	return svg
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r > 0 {
		return math.Sin(r) / r
	}

	return 0.0
}
