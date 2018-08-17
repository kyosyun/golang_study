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

func TestCheckAnagram2(t *testing.T) {
	for _, test := range []struct{
		a, b string
		exected bool
	}{
		{"a","a",true},
		{"a","b",false},
		{"日本","日本",true},
		{"日本","本日",true},
		{"日本日本","日本",false},
		{"日a本","日b本",false},
	}{
		result := anagrams.CheckAnagram(test.a, test.b)
		if result != test.exected {
			t.Errorf("result = %v, but was = %vk", result, test.exected)
		}
	}
}