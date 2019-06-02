package main

import (
	"testing"
)

func TestCaclScore(t *testing.T) {
	tests := []struct {
		testScore int
		absentNum int
		expected  int
	}{
		{100, 0, 100},
		{100, 2, 90},
		{100, 20, 0},
		{100, 30, 0},
	}

	for _, test := range tests {
		res := calcScore(test.testScore, test.absentNum)
		if test.expected != res {
			t.Errorf("result is %v, but expected is %v", res, test.expected)
		}
	}
}

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
