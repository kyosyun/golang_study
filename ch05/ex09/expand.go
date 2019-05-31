package ex09

import "regexp"

// $foo をf("foo")が返すテキストでｄ置換する関数を返却する
func Expand(s string, f func(string) string) string {
	pattern := regexp.MustCompile(`\$\w+`)
	wrapper := func(arg string) string {
		//最初の$を取り除く
		return f(arg[1:])
	}
	return pattern.ReplaceAllStringFunc(s, wrapper)
}
