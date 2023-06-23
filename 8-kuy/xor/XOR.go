package xor

import (
	"math"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if i == n {
			return true
		}
		if i*i <= n {
			if n%i == 0 {
				return false
			}
			/* if i > int(math.Log(float64(n))) {
				return true
			} */
			if i > int(math.Sqrt(float64(n))) {
				return true
			}
		}
	}
	return true
}
