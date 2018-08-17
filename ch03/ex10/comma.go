package comma

import "bytes"

func Comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func CommaWithBuffer(s string) string {

	buf := bytes.Buffer{}
	runes := []byte(s)
	count := len(runes) % 3
	for _, c := range runes {
		if count == 0 {
			buf.WriteRune(',')
			count = 3
		}
		buf.WriteByte(c)
		count--
	}
	return buf.String()
}