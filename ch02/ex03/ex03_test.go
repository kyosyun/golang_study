package ex03_test

import (
"testing"
"golang_study/ch02/ex03"
)


func TestPopcount(t *testing.T) {
	i := ex03.PopCount(255)
	if (i == 8) {
		t.Log("success")
	} else {
		t.Error("err when 255")
	}

	i = ex03.PopCount(3)
	if (i == 2) {
		t.Log("success")
	} else {
		t.Error("err when 255")
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		ex03.PopCount(255)
	}
}