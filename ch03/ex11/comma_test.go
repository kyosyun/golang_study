package comma_test

import (
	"testing"
	"golang_study/ch03/ex11"
)

func TestComma(t *testing.T) {
	print(comma.CommaWithSinged("1235") + "\n")
	print(comma.CommaWithSinged("+1235") + "\n")
	print(comma.CommaWithSinged("+1235.12345") + "\n")
	print(comma.CommaWithSinged("-1235.12345") + "\n")
	print(comma.CommaWithSinged("+12351234.5") + "\n")
	print(comma.CommaWithSinged("-12351234.5") + "\n")
}

func TestCommaWithSinged(t *testing.T) {
	for _, test := range []struct {
		data string
		expected string
	}{
		//{"",""},
		{"123","123"},
		{"1234","1234"},
		{"1.234","1.234"},
		{"12.34","12.34"},
		{"123.4","123.4"},

		{"-123","-123"},
		{"-1234","-1234"},
		{"-1.234","-1.234"},
		{"-12.34","-12.34"},
		{"-123.4","-123.4"},

		{"+123","+123"},
		{"+1234","+1234"},
		{"+1.234","+1.234"},
		{"+12.34","+12.34"},
		{"+123.4","+123.4"},

	} {
		result := comma.CommaWithSinged(test.data)
		if result != test.expected {
			t.Errorf("result = %q, but want %q", result, test.expected)
		}
	}
}