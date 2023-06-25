package twosum

import "sort"

// задача ускорить алгоритм
func TwoSum(numbers []int, target int) [2]int {
	sort.Slice(numbers, func(i2, j2 int) bool {
		return numbers[i2] < numbers[j2]
	})
	res := [2]int{}
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			res[0] = i
			res[1] = j
			break
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return res
}
