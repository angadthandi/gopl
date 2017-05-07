package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ex18checkPrefix(uptr *string) {
	b := strings.HasPrefix(*uptr, "http://")
	if !b {
		b = strings.HasPrefix(*uptr, "https://")
		if !b {
			*uptr = "http://" + *uptr
		}
	}
}

func Fetchurls() {
	for _, url := range os.Args[1:] {
		ex18checkPrefix(&url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
		fmt.Printf("Ex18 status: %s\n", resp.Status)
	}
}

func Ex17Fetchurl() {
	for _, url := range os.Args[1:] {
		ex18checkPrefix(&url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("Ex18 status: %s\n", resp.Status)
	}
}
