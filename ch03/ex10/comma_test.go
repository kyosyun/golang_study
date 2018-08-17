package comma_test

import (
	"testing"
	"golang_study/ch03/ex10"
)

func TestComma(t *testing.T) {
	print(comma.Comma("1235") + "\n")
}

func TestCommaWithBuffer(t *testing.T) {
	print(comma.CommaWithBuffer("1235") + "\n")
}