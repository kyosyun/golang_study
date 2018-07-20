package ex05_test

import (
	"golang_study/ch02/ex03"
	"golang_study/ch02/ex04"
	"golang_study/ch02/ex05"
	"golang_study/ch02/popcount"
	"testing"
)

func TestPopcount(t *testing.T) {
	for i := 0; i < 256; i++ {
		if ex05.PopCount(uint64(i)) == ex04.PopCount(uint64(i)) {
			t.Log("success: ", i)
		} else {
			t.Error("fail: when:", i, "i: ", ex05.PopCount(uint64(i)), "expected: ", ex04.PopCount(uint64(i)))
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// ロジックによって0や1の処理のスピードが偏るので、PopCountの引数をばらつかせる
		popcount.PopCount(uint64(b.N % 255))
	}
}

func BenchmarkPopCountEx03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// ロジックによって0や1の処理のスピードが偏るので、PopCountの引数をばらつかせる
		ex03.PopCount(uint64(b.N % 255))
	}
}

func BenchmarkPopCountEx04(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex04.PopCount(uint64(b.N % 255))
	}
}

func BenchmarkPopCountEx05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex05.PopCount(uint64(b.N % 255))
	}
}
