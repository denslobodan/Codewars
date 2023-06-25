package fibonacci

func ProductFib(prod uint64) [3]uint64 {
	// your code
	var cache1, cache2, cache3 uint64 = 1, 1, 0

	res := 0
	for i := 2; i <= int(prod); i++ {
		cache3 = cache1 + cache2
		cache1, cache2 = cache2, cache3
		if cache1*cache2 == prod {
			res = 1
			break
		} else if cache1*cache2 > prod {
			break
		}
	}
	return [3]uint64{cache1, cache2, uint64(res)}
}
