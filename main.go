package main

import (
	"codewars/6-kuy/walk"
	"fmt"
)

func main() {
	w := walk.IsValidWalk([]rune{'s', 'w', 'w', 's', 'e', 's', 's', 'w', 'e', 'w'})
	alW := walk.AlternativeWalk([]rune{'s', 'w', 'w', 's', 'e', 's', 's', 'w', 'e', 'w'})
	fmt.Println(w, alW)
}
