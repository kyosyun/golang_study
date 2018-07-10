// 入力に2回以上現れた行の数とその行のテキストを表示する
// 標準入力から読み込むか名前が指定されたファイルの一覧から読み込む
// 重複した行のそれぞれが含まれていたファイルの名前も表示する
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countedFile := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countedFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countedFile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t files: %s\n", n, line, countedFile[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCounts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if !contains(f.Name(), fileCounts[input.Text()]) {
			fileCounts[input.Text()] = append(fileCounts[input.Text()], f.Name())
		}
		counts[input.Text()]++
	}
}

func contains(target string, list []string) bool {
	for _, value := range list {
		if target == value {
			return true
		}
	}
	return false
}
