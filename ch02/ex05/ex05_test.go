package ex05_test

import (
"fmt"
"testing"
"golang_study/ch02/ex05"
	"golang_study/ch02/ex03"
	"golang_study/ch02/ex04"
)

func TestPopcount(t *testing.T) {
	i := ex05.PopCount(255)
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

func BenchmarkPopCountEx05(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		ex05.PopCount(255)
	}
}