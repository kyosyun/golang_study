package ex02_test

import (
	"testing"
	"golang_study/ch01/ex02"
)

func TestEx02(t *testing.T) {
	ex02.Join([]string{"a","b","c"})
}