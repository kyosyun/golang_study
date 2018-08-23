package main

import (
	"fmt"
	"os"

	"strings"

	"golang.org/x/net/html"
)

func main() {
	//標準入力の読み込み
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "extractText: %v\n", err)
	}

	//標準入力の内容を処理して、tagsにタグのカウントを格納する
	printExtractText(nil, doc)

}

//htmlNode: 最初は標準入力の内容を受け取る。
func printExtractText(stack []string, n *html.Node) {

	switch n.Type {
	case html.ElementNode:
		stack = append(stack, n.Data)
	case html.TextNode:
		lastTag := stack[len(stack)-1]
		if lastTag != "script" && lastTag != "style" {
			if len(strings.TrimSpace(n.Data)) > 0 {
				fmt.Printf("<%s>%s</%s>\n", lastTag, n.Data, lastTag)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printExtractText(stack, c)
	}
}
