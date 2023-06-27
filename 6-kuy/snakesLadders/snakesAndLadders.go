package snakesladders

func SnakesAndLadders(board, dice []int) int {
	// Your code here
	current := 0
	for _, val := range dice {
		if current+val >= len(board) {
			continue
		}
		current += val

		if board[current] != 0 {
			current += board[current]
		}
		if current == len(board)-1 {
			break
		}
	}
	return current
}
