package ex03_test

import (
	"testing"
	"golang_study/ch01/ex03"
)

func TestEx03(t *testing.T) {
	ex03.Join([]string{"a","b","c"})
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex03.Join([]string{"a","b","c"})
	}
}