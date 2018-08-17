package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

func main() {

	//入力を受け付ける
	reader := bufio.NewReader(os.Stdin)

	for {
		bytes, err := reader.ReadBytes('\n')

		if err == io.EOF {
			os.Exit(0)
		}

		if err != nil {
			fmt.Printf("read failed: %v", err)
			os.Exit(1)
		}

		//入力に対して暗号化する
		sumResult256 := sha256.Sum256(bytes)
		sumResult384 := sha512.Sum384(bytes)
		sumResult512 := sha512.Sum512(bytes)

		//出力する
		fmt.Printf("sum256:%x\n", sumResult256)
		fmt.Printf("sum384:%x\n", sumResult384)
		fmt.Printf("sum512:%x\n", sumResult512)
	}
}
