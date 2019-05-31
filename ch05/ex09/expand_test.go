package ex09_test

import (
	"golang_study/ch05/ex09"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		res     string
		expect  string
		wrapper func(s string) string
	}{
		{
			"hoge $hoge hoge",
			"hoge HOGE hoge",
			func(s string) string { return strings.ToUpper(s) },
		},
		{
			"$FUGA FUGA $FUGA",
			"fuga FUGA fuga",
			func(s string) string { return strings.ToLower(s) },
		},
	}

	for _, test := range tests {
		res := ex09.Expand(test.expect, test.wrapper)
		if test.expect != res {
			t.Errorf("result is %s, but expected is %s", res, test.expect)
		}
	}
}
