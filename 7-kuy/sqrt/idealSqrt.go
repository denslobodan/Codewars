package sqrt

import (
	"fmt"
	"math"
)

func FindNextSquare(sq int64) int64 {
	nextSq := 0.0
	multi := math.Sqrt(float64(sq))
	if sq%int64(multi) != 0 {
		defer fmt.Printf("since %d is not a perfect square", sq)
	} else {
		nextMulti := multi + 1
		nextSq = nextMulti * nextMulti
		return int64(nextSq)
	}
	return -1
}
