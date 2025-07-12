package main

func placeWord(board [][]string, word string, row, col int, vertical bool) {
	if vertical {
		for i, ch := range word {
			board[row+i][col] = string(ch)
		}
	} else {
		for i, ch := range word {
			board[row][col+i] = string(ch)
		}
	}
}

func canPlaceWord(board [][]string, word string, row, col int, vertical bool) bool {
	if vertical {
		if row+len([]rune(word)) > len(board) {
			return false
		}
		for i, ch := range word {
			cell := board[row+i][col]
			if cell != " " && cell != string(ch) {
				return false
			}
		}
	} else {
		if col+len([]rune(word)) > len(board[0]) {
			return false
		}
		for i, ch := range word {
			cell := board[row][col+i]
			if cell != " " && cell != string(ch) {
				return false
			}
		}
	}
	return true
}
