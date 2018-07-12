package ex04_test

import (
"fmt"
"testing"
"golang_study/ch02/ex04"
)

func TestPopcount(t *testing.T) {
	i := ex04.PopCount(1)
	fmt.Println(i)
}