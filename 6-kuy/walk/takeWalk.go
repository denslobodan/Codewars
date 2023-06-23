package walk

func IsValidWalk(walk []rune) bool {
	success := false
	if len(walk) != 10 {
		return success
	}
	m := make(map[rune]int)
	for _, v := range walk {
		m[v] += 1
	}
	if m['n'] != m['s'] || m['w'] != m['e'] {
		return success
	} else {
		success = true
	}
	return success
}

// альтернативное решение
func AlternativeWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	res := 0
	for _, v := range walk {
		switch v {
		case 's':
			res--
		case 'n':
			res++
		case 'w':
			res++
		default:
			res--
		}
	}
	return res == 0
}

/* x, y := 0, 0
for _, r := range walk {
  switch r {
  case 'n': y++
  case 's': y--
  case 'e': x++
  case 'w': x--
  }
} */
