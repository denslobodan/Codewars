package romanNumerals

func Decode(roman string) int {
	res := 0
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(roman)-1; i++ {
		if romanMap[roman[i]] < romanMap[roman[i+1]] {
			res -= romanMap[roman[i]]
		} else if romanMap[roman[i]] >= romanMap[roman[i+1]] {
			res += romanMap[roman[i]]
		}
	}
	res = res + romanMap[roman[len(roman)-1]]
	return res
}
