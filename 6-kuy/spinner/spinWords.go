package spinner

import (
	"bytes"
	"strings"
)

func SpinWords(str string) string {
	strList := strings.Split(str, " ")
	buffer := bytes.Buffer{}
	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}
	for i, s := range strList {
		if len(s) < 5 {
			buffer.WriteString(s)
		} else {
			buffer.WriteString(reverse(s))
		}
		if i < len(strList)-1 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
} // SpinWords
