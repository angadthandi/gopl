package main

import "fmt"
import "net/http"

// import "github.com/angadthandi/gopl/ch1"
// import "github.com/angadthandi/gopl/ch2"
import "github.com/angadthandi/gopl/ch3"

func main() {
	// ch1.Echo1()
	// ch1.Echo2()
	// ch1.Echo3()
	// ch1.Dup1()
	// ch1.Dup2()
	// ch1.Dup3()
	// ch1.Fetchurls()
	// ch1.Ex17Fetchurl()
	// ch1.FetchAll()
	// ch1.Server1()
	// ch1.Server2()
	// ch1.Server3()
	// ch1.Server4()

	// ch2.Echo4()
	// ch2.Tempconv()
	// ch2.Cf()
	// ch2.GeneralConv()

	// ch3.Surface()

	http.HandleFunc("/", handle)
	http.ListenAndServe(":9000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	// ch3.Surface(w)
	svg := ch3.Surface()
	fmt.Fprintf(w, "%s", svg)
}
