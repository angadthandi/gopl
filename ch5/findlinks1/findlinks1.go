// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var linkNodes = map[string]bool{
	"a":      true,
	"img":    true,
	"script": true,
	"link":   true,
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	// for _, link := range visit(nil, doc) {
	// 	fmt.Println(link)
	// }
	visit(nil, doc)
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	// if n.Type == html.ElementNode && n.Data == "a" {
	if n.Type == html.ElementNode && linkNodes[n.Data] {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
				fmt.Println(links)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
