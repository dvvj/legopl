package links

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ExtractBase(reader io.Reader) ([]string, error) {
	doc, err := html.Parse(reader)

	if err != nil {
		return nil, fmt.Errorf("parsing as HTML: %v", err)
	}

	var links []string

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				// link, err := resp.Request.URL.Parse(a.Val)
				// if err != nil {
				// 	continue
				// }
				// links = append(links, link.String())
				links = append(links, a.Val)
			}
		}
	}

	forEachNode(doc, visitNode, nil)

	return links, nil
}

func ExtractUrl(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	res, err := ExtractBase(resp.Body)
	resp.Body.Close()

	var fixedLinks []string
	for _, link := range res {
		fixed, err := resp.Request.URL.Parse(link)
		if err != nil {
			continue
		}
		fixedLinks = append(fixedLinks, fixed.String())
	}
	return fixedLinks, err
}

func ExtractFile(file string) ([]string, error) {
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		return nil, fmt.Errorf("cannot open file %s", file)
	}

	reader := bufio.NewReader(f)

	res, err := ExtractBase(reader)
	return res, err
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
