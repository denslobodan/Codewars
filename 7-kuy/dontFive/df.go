package dontfive

import (
	"strconv"
	"strings"
)

func DontGiveMeFive(start int, end int) int {
	sl := []int{}
	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if strings.Contains(s, "5") {
			continue
		} else {
			sl = append(sl, i)
		}
	}
	return len(sl)
}
