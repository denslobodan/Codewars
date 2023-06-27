package ministingfuck

func Interpreter(code string) string {
	memory := make([]int, 60000)
	ptr := 0
	s := ""
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			memory[ptr]++
		case '.':
			s += string(rune(memory[ptr]))
		}
	}
	return s
}
