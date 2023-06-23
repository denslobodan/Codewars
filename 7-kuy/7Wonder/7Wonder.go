package wonder

import "math"

// Calculate how many points you would make in 7 Wonders from these science cards!
func SevenWondersScience(compasses int, gears int, tablets int) int {
	// TODO: Do the code here
	var glyphs = 7
	sl := make([]int, 0)
	sl = append(sl, compasses, gears, tablets)
	total := 0
	min := math.Min(math.Min(float64(compasses), float64(gears)), float64(tablets))
	total = int(min) * glyphs
	for _, v := range sl {
		v *= v
		total += v
	}
	return total
}
