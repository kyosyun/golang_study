package comma

import (
	"strings"
)

func CommaWithSinged(s string) string{
	//カンマチェック
	if s[0:1] == "+" || s[0:1] == "-" {
		return s[0:1] + commaWithFlotingNum(s[1:])
	}

	return commaWithFlotingNum(s)
}

func addComma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return addComma(s[:n-3]) + "," + s[n-3:]
}

func commaWithFlotingNum(s string) string {
	//小数点のチェック
	pIndex := strings.IndexByte(s, '.')
	if pIndex >= 0 {
		return addComma(s[:pIndex]) + s[pIndex:]
	}
	return addComma(s)
}