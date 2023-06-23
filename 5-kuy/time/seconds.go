package mytime

import (
	"fmt"
)

func HumanReadableTime(seconds int) string {
	// your code here
	if seconds < 0 || seconds > 359999 {
		return "non-format"
	}
	var h, m, s int
	h = seconds / 3600
	m = (seconds - h*3600) / 60
	s = seconds - h*3600 - m*60

	return fmt.Sprintf("%.2d:%.2d:%.2d", h, m, s)
}
