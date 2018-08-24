package main

import (
	"fmt"
	"net/http"

	"strings"

	"os"

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
		wl := strings.Split(n.Data, " ")
		words += len(wl)
		for _, w := range wl {
			if len(strings.TrimSpace(w)) > 0 {
				//デバッグ用　wordを出力する
				println(strings.TrimSpace(w))
				words++
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		images += i
		words += w
	}
	return words, images
}
