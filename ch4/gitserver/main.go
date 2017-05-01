package main

import (
	"github.com/angadthandi/gopl/ch4/github"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	renderIssues(w)
}

func renderIssues(out http.ResponseWriter) {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var tmpl = template.Must(template.New("issueList").
		ParseFiles("issueList.go.html"))
	tmpl.ExecuteTemplate(out, "issueList.go.html", result)

	// t, _ := template.ParseFiles("issueList.html")
	// t.Execute(out, result)
}
