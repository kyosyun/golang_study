package comma_test

import (
	"testing"
	"golang_study/ch03/ex10"
	"fmt"
)

func TestComma(t *testing.T) {
	print(comma.Comma("1235") + "\n")
}

func TestCommaWithBuffer(t *testing.T) {
	print(comma.CommaWithBuffer("1235") + "\n")

	for _, test := range []struct{
		data string
		expected string
	}{
		{"", ""},
		{"123","123"},
		{"123","123"},
		{"123456","123,456"},
		{"123456789","123,456,789"},
	}{
		result := comma.Comma(test.data)
		if result != test.expected {
			fmt.Printf("expected: %s, result: %s",test.expected, result)
		}
	}
}