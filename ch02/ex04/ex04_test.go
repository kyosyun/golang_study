package ex04_test

import (
"fmt"
"testing"
"golang_study/ch02/ex04"
	"golang_study/ch02/ex03"
)

func TestPopcount(t *testing.T) {
	i := ex04.PopCount(1)
	fmt.Println(i)
}

func BenchmarkPopCountEx03(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		ex03.PopCount(255)
	}
}

func BenchmarkPopCountEx04(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		ex04.PopCount(255)
	}
}