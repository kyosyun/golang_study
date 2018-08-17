package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	//標準入力の読み込み
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}

	//標準入力の内容を処理するして、linkの配列を作成し、printを行う。
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//htmlNode: 最初は標準入力の内容を受け取る。
func visit(links []string, n *html.Node) []string {

	//
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
