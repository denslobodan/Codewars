package main

import (
	snakesladders "codewars/6-kuy/snakesLadders"
	"fmt"
)

func main() {
	board := []int{0, 2, 0, 0, 0, 2, 5, 0, -6, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, -9, 0, 0, 0, 0, 0, -4, 0, -4, 0, -3}
	rolls := []int{1, 2, 5, 4, 1, 1, 5, 4, 2, 1}

	res := snakesladders.SnakesAndLadders(board, rolls)
	fmt.Println(res)
}
