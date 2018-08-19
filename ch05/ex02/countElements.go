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

	// 結果を格納するmap
	tags := make(map[string]int)

	//標準入力の内容を処理して、tagsにタグのカウントを格納する
	visit(tags, doc)

	for key, _ := range tags {
		fmt.Printf("%s = %d\n", key, tags[key])
	}
}

//htmlNode: 最初は標準入力の内容を受け取る。
func visit(tags map[string]int, n *html.Node) {

	if n.Type == html.ElementNode {
		tags[n.Data] = tags[n.Data] + 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(tags, c)
	}
}
