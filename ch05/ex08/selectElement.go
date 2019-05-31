package main

import (
	"fmt"

	"net/http"
	"os"

	"io"

	"golang.org/x/net/html"
)

var depth int

//[ ] pre関数にてtrueを返却するようにする
//[ ] post関数にてtrueを返却するようにする
//[ ] 指定したIDをのHTML要素を見つけた場合には、操作を中止して、該当の要素を出力する

// テストコードを書くために指定
var writer io.Writer

func main() {

	if len(os.Args) != 3 {
		os.Exit(1)
	}
	err := selectElementById(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("err:%v", err)
	}

}

// urlのなかで、IDを保持している最初のElementを返却する
func selectElementById(url, id string) error {
	// 出力先
	writer = os.Stdout
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return err
	}

	findNode := ElementById(doc, id)
	printNode(findNode)
	return nil
}

func ElementById(n *html.Node, id string) *html.Node {
	return forEachNode(n, id, startElement(id), nil)
}

func forEachNode(n *html.Node, id string, pre func(n *html.Node) bool, post func(n *html.Node) *html.Node) *html.Node {
	found := false
	if pre != nil {
		found = pre(n)
	}

	if found {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		post(n)
	}
	return nil
}

func startElement(id string) func(n *html.Node) bool {
	return func(n *html.Node) (found bool) {
		if n.Type == html.ElementNode {
			//要素を出力する
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return true
				}
			}

		}
		return
	}
}

func printNode(n *html.Node) {
	if n == nil {
		fmt.Fprintf(writer, "cant find id")
		return
	}
	fmt.Fprintf(writer, "<%s", n.Data)
	for _, a := range n.Attr {
		fmt.Fprintf(writer, " %s=\"%s\"", a.Key, a.Val)
	}
	fmt.Fprintf(writer, ">")
}
