package pretty

import (
	"fmt"

	"net/http"
	"os"

	"strings"

	"io"

	"golang.org/x/net/html"
)

var depth int

//[x]コメントノードを表示する
//[x]テキストノードを表示する
//[x]要素の属性を表示する
//[x]要素が子を持たない場合には、一行でタグを閉じる
//[]テストコードを書く

// テストコードを書くために指定
var writer io.Writer

//func main() {
//	for i := 0; i < len(os.Args); i++ {
//		prettyPrint(os.Args[i])
//	}
//}

// urlへGETリクエストを行い、ドキュメント内に含まれる単語と画像の数を返す。
func prettyPrint(url string) {
	writer = os.Stdout
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, hasChild bool)) {

	//要素が子を持つか持つか持たないか判定
	hasChild := false
	if n.FirstChild != nil {
		hasChild = true
	}

	//子を持つ場合には、子に対して再帰呼び出しさせる -> postを呼び出す。
	if pre != nil {
		pre(n, hasChild)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n, hasChild)
	}
}

func startElement(n *html.Node, hasChild bool) {
	switch n.Type {
	case html.DoctypeNode:
		fmt.Fprintf(writer, "<!DOCTYPE %s>\n", n.Data)
	case html.ElementNode:
		//インデントしてタグの開始
		fmt.Fprintf(writer, "%*s<%s", depth*2, "", n.Data)
		//要素を出力する
		for _, a := range n.Attr {
			fmt.Fprintf(writer, " %s=\"%s\"", a.Key, a.Val)
		}
		//タグの終了を出力
		if hasChild {
			fmt.Fprint(writer, ">\n")
			depth++
		} else {
			fmt.Fprintf(writer, "/>\n")
		}
	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) > 0 {
			fmt.Fprintf(writer, "%*s%s\n", depth*2, "", n.Data)
		}
	case html.CommentNode:
		fmt.Fprintf(writer, "%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node, hasChild bool) {
	if n.Type == html.ElementNode {
		if hasChild {
			depth--
			fmt.Fprintf(writer, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
