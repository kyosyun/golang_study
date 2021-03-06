package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
)

func main() {
	start := time.Now()
	//チャンネルの定義
	ch := make(chan string)
	ch2 := make(chan string)
	for _,url := range os.Args[1:] {
		// ゴルーチンの開始
		go fetch(url, ch, 1)
		go fetch(url, ch2, 2)
	}
	for range os.Args[1:] {
		// chチャンネルから受信する
		fmt.Println(<-ch)
		fmt.Println(<-ch2)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string, index int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//出力用のファイルの作成
	dst, err := os.Create(fmt.Sprintf("out_%d.html", index))
	if err != nil {
		fmt.Fprint(os.Stderr, "fetch: reading %s: %v \n", url, err)
		os.Exit(1)
	}
	defer dst.Close()

	// ioutil.Discard: ファイルから読み込んで即座に破棄する。返り値のbytesのみ利用する
	nbytes, err := io.Copy(dst, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("While reading %s: %v ", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}