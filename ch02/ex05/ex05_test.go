package ex05_test

import (
"fmt"
"testing"
"golang_study/ch02/ex05"
)

func TestPopcount(t *testing.T) {
	i := ex05.PopCount(1)
	fmt.Println(i)
}