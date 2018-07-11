package ex10

import (
"time"
"os"
"fmt"
"net/http"
"io"
"io/ioutil"
)

func main() {
	start := time.Now()
	//チャンネルの定義
	ch := make(chan string)

	//出力用のファイルの作成
	dst, err := os.Create("out.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, "fetch: reading %s: %v \n", url, err)
		os.Exit(1)
	}
	defer dst.Close()

	for _,url := range os.Args[1:] {
		// ごルーチンの開始
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// chチャンネルから受信する
		_, err = io.Copy(dst, resp.Body)

		dst.WriteString(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// ioutil.Discard: ファイルから読み込んで即座に破棄する。返り値のbytesのみ利用する
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("While reading %s: %v ", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}