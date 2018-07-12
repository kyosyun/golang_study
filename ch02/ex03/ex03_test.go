package ex03_test

import (
"fmt"
"testing"
"golang_study/ch02/ex03"
)

func TestPopcount(t *testing.T) {
	i := ex03.PopCount(1)
	fmt.Println(i)
}