package ministingfuck

import (
	"bufio"
	"os"
)

func Interpreter(code string) string {
	var memory = make([]int, 1000)
	var ptr int
	var output string

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case ',':
			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadRune()
			if err != nil {
				// handle error
			}
			memory[ptr] = int(char)
		case '.':
			output += string(rune(memory[ptr]))
		case '>':
			ptr++
		case '<':
			ptr--
		case '[':
			if memory[ptr] == 0 {
				for j := i; j < len(code); j++ {
					if code[j] == ']' {
						i = j
						break
					}
				}
			}
		case ']':
			if memory[ptr] != 0 {
				for j := i; j >= 0; j-- {
					if code[j] == '[' {
						i = j
						break
					}
				}
			}
		}
	}
	return output
}
