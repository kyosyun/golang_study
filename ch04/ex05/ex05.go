// []stringスライス内で隣接している重複をスライス内で除去する関数を書く
package ex05

func deleteDuplicate(str []string) []string {
	if len(str) == 0 {
		return str
	}

	current := 0
	for i := 0; i < len(str)-1; i++ {
		if str[current] != str[i+1] {
			str[current+1] = str[i+1]
			current++
			continue
		}
	}

	return str[:current+1]
}