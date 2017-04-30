package ch4

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	api = "https://www.omdbapi.com/?"
	dst = "images/"
)

type ApiResponse struct {
	Response string
	Error    string
	Title    string
	Poster   string
}

func Poster() {
	var response ApiResponse
	var t string
	if len(os.Args) > 1 {
		t = os.Args[1]
	}

	url := api + "t=" + t
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("GET request failed: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("search query failed: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Errorf("Json Decode failed: %s", err)
	}

	if response.Response == "False" {
		fmt.Printf("Error = %v\n", response.Error)
	}

	if response.Response == "True" {
		fmt.Printf("Title = %v\n", response.Title)
		fmt.Printf("Poster = %v\n", response.Poster)
		download(response.Title, response.Poster)
	}

	resp.Body.Close()
}

func download(t, f string) {
	fTokens := strings.Split(f, ".")
	ext := fTokens[len(fTokens)-1]

	path, err := filepath.Abs(dst + t + "." + ext)
	if err != nil {
		fmt.Printf("Error filepath : %v\n", err)
		return
	}

	out, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file : %v\n", err)
		return
	}
	defer out.Close()

	resp, err := http.Get(f)
	if err != nil {
		fmt.Printf("Error downloading poster : %v\n", err)
		return
	}
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error downloading poster : %v\n", err)
		return
	}

	fmt.Println(n, " bytes downloaded!")
}
