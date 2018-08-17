//一回のパス操作でrotateを書く
package main

import "fmt"

func rotate(data []int, rotateSize int) []int {
	// rotateSizeの分rotateする
	for i := 0; i < len(data)-rotateSize; i++ {
		data[i] = data[i+rotateSize]
	}
	return data[:len(data)-rotateSize]
}

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(a)
	rotate(a,3)
	fmt.Println(a)
}