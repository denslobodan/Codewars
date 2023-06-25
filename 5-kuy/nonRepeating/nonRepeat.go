package nonrepeating

import "strings"

func FirstNonRepeating(str string) string {
	var s string
	m := make(map[rune][]int, 0)
	strLow := strings.ToLower(str)
	for i, v := range strLow {
		m[v] = append(m[v], i)
	}
	for i, v := range strLow {
		if len(m[v]) == 1 {
			s = string(str[i])
			break
		}
	}
	return s
}
