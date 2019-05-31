package main

import (
	"time"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

// https://xkcd.com/571/info.0.json
// 複数のコミックのURLをダウンロードする
//  - https://xkcd.com/571/info.0.jsonのようにダウンロードする

// ダウンロードした内容を元に、オフラインのインデックスをつくる
// コマンドラインで提供された検索後と一致するコミックのURLと内容を表示する
func main() {

}

func getCommic() {
	start := time.Now()
	//チャンネルの定義
	ch := make(chan string)
	for i := 1 ; i < 10; i++  {
		// ゴルーチンの開始
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)

		// データの取得
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}

		// ioutil.Discard: ファイルから読み込んで即座に破棄する。返り値のbytesのみ利用する
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()


	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}