package main

import (
	"fmt"
	// "github.com/angadthandi/gopl/ch1"
	// "github.com/angadthandi/gopl/ch2"
	"github.com/angadthandi/gopl/ch3"
	"github.com/angadthandi/gopl/ch4"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// ch1.Echo1()
	// ch1.Echo2()
	// ch1.Echo3()
	// ch1.Dup1()
	// ch1.Dup2()
	// ch1.Dup3()
	// ch1.Fetchurls()
	// ch1.Ex17Fetchurl()
	// ch1.FetchAll()	// ch1.Server1()
	// ch1.Server2()
	// ch1.Server3()
	// ch1.Server4()

	// ch2.Echo4()
	// ch2.Tempconv()
	// ch2.Cf()
	// ch2.GeneralConv()

	// ch3.Surface()

	// http.HandleFunc("/", handle)
	// http.HandleFunc("/", WebHandler)
	// http.ListenAndServe(":9000", nil)

	// ch3.Comma()
	// ch3.Anagram()

	// ch4.ShaDiffBits()
	// ch4.ShaOut()
	// ch4.Reverse()
	// ch4.Charcount()
	// ch4.Wordfreq()
	// ch4.Issues()
	// ch4.Xkcd()
	ch4.Poster()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Go!")

	w.Header().Set("Content-Type", "image/svg+xml")
	svg := ch3.Surface()
	fmt.Fprintf(w, "%s", svg)

	// w.Header().Set("Content-Type", "image/png")
	// png := ch3.Mandelbrot()
	// ret := http.FileServer(http.Dir(png))
	// fmt.Fprintf(w, "%s", ret)
	// ch3.Mandelbrot(w)
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	var Path = ch3.Mandelbrot()
	// ResizeImage(Path, 500)

	img, err := os.Open(Path)
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
	io.Copy(w, img)
}
