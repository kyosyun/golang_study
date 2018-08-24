package main

import (
	"fmt"
	"net/http"

	"os"

	"bufio"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(CountWordsAndImages(os.Args[i]))
	}
}

// urlへGETリクエストを行い、ドキュメント内に含まれる単語と画像の数を返す。
func CountWordsAndImages(url string) (words, images int, err error) {
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
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	//画像をカウントする: タグがimgのものをカウントする
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
		//デバッグ用 ImageSrcを出力する
		//for _, a := range n.Attr {
		//	if a.Key == "src" {
		//		println(a.Val)
		//	}
		//}
	}

	//単語数をカウントする。
	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			//println(input.Text())
			words++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return words, images
}
