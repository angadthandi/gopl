package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementMaps: %v\n", err)
		os.Exit(1)
	}

	test := make(map[string]int)
	for el, c := range elemMap(test, doc) {
		fmt.Printf("%v = %v\n", el, c)
	}

	for _, v := range textNodes(nil, doc) {
		fmt.Printf("TextNodeData = %v\n", v)
	}
}

func elemMap(elMaps map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elMaps[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elMaps = elemMap(elMaps, c)
	}
	return elMaps
}

func textNodes(stack []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Data != "style" && n.Data != "script" {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		textNodes(stack, c)
	}

	return stack
}
