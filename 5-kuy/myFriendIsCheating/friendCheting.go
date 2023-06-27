package myfriendischeating

/*
	нахождения двух чисел в последовательности от 1 до n,

произведение которых равно сумме всех чисел в последовательности,
исключая 2 выбранных числа
*/

func RemovNb(n uint64) [][2]uint64 {
	sum := n * (n + 1) / 2
	result := make([][2]uint64, 0)
	result = nil
	for a := uint64(1); a <= n; a++ {
		b := (sum - a) / (a + 1)
		if b <= n && a*b == sum-a-b {
			result = append(result, [2]uint64{a, b})
		}
	}

	return result
}
