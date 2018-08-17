package anagrams_test

import (
	"testing"
	"golang_study/ch03/ex12"
)

func TestCheckAnagram(t *testing.T) {
	println(anagrams.CheckAnagram("abc", "abc"))
	println(anagrams.CheckAnagram("abc", "cba"))
	println(anagrams.CheckAnagram("abc", "abd"))
}
