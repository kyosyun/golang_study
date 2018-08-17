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
	visit(doc)
}

//htmlNode: 最初は標準入力の内容を受け取る。
func visit(n *html.Node) {

	//
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
