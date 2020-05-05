package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc, 0)
}

func outline(stack []string, n *html.Node, indent int) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		idt := strings.Repeat("  ", indent)
		fmt.Print(idt)
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c, indent+1)
	}
}
