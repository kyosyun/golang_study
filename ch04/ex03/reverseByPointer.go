// スライスの代わりに配列へのポインターを使うようにreverseを書き直す
package reverseByPointer

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

const Size = 10
func reverseByPointer(a *[Size]int, len int) {
	for i, j := 0, Size-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}