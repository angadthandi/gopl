package ch4

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const indexFilePath = "ch4/xkcd_index.txt"

type IndexSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Indexdata
}

type Indexdata struct {
	Num        int
	HTMLURL    string `json:"url"`
	Title      string
	Transcript string
	Day        string
	Month      string
	Year       string
	ImageURL   string
}

func Xkcd() {
	var result IndexSearchResult
	var q string
	if len(os.Args) > 1 {
		q = os.Args[1]
	}

	file, err := os.Open(indexFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err = %v\n", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&result.Items)
	if err != nil {
		fmt.Printf("err = %v\n", err)
		return
	}

	for _, v := range result.Items {
		if q == "" || strings.Contains(v.HTMLURL, q) || strings.Contains(v.Transcript, q) {
			fmt.Println("======================\n")
			fmt.Printf("HTMLURL = %v\n", v.HTMLURL)
			fmt.Printf("Transcript = %v\n", v.Transcript)
		}
	}
}
