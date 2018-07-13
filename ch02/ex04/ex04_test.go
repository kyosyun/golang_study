package ex04_test

import (
"testing"
"golang_study/ch02/ex04"
	"golang_study/ch02/ex03"
)

func TestPopcount(t *testing.T) {
	for i := 0; i < 256 ; i++  {
		if(ex03.PopCount(uint64(i)) == ex04.PopCount(uint64(i))) {
			t.Log("success: ", i)
		} else {
			t.Error("fail: when:", i, "i: ", ex04.PopCount(uint64(i)), "expected: ", ex03.PopCount(uint64(i)))
		}
	}
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