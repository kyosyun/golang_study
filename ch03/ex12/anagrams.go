package anagrams

import "strings"

func CheckAnagram(s1 string, s2 string) bool{
	//s1, s2 の長さチェック
	lenS1 := len(s1)
	lenS2 := len(s2)

	if lenS1 != lenS2 {
		return false
	}

	//s1に含まれている文字を一文字ずつ、s2から消していく
	runesS1 := []byte(s1)
	for _, r := range runesS1 {
		index := strings.IndexByte(s2, r)
		if index >= 0 {
			s2 = s2[:index] + s2[index+1:]
			continue
		} else {
			return false
		}
	}
	return true
}

