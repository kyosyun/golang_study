//一回のパス操作でrotateを書く
package ex04

func rotate(data []int, rotateSize int) []int {
	// rotateSizeの分rotateする
	for i := 0; i < len(data)-rotateSize; i++ {
		data[i] = data[i+4]
	}

	return data[:len(data)-rotateSize]
}