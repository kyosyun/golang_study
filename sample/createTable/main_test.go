package main

import (
	"testing"
)

func TestMethod(t *testing.T) {
	tests := []struct {
		in     string
		expect string
	}{
		{
			"hoge $hoge hoge",
			"hoge HOGE hoge",
		},
	}
	for _, test := range tests {
		res := Mod(test.in)
		if test.expect != res {
			t.Errorf("result is %s, but expected is %s", res, test.expect)
		}
	}
}

func Mod(mod string) string {
	return mod
}
