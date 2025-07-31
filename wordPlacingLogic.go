package main

import "strings"

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
	word = strings.ToLower(word) // Ensure consistent case

	if vertical {
		// Check if word fits vertically
		if row+len(word) > len(board) {
			return false
		}
		// Check each position
		for i, ch := range word {
			if row+i >= len(board) || col >= len(board[0]) {
				return false
			}
			cell := board[row+i][col]
			if cell != " " && strings.ToLower(cell) != string(ch) {
				return false
			}
		}
	} else {
		// Check if word fits horizontally
		if col+len(word) > len(board[0]) {
			return false
		}
		// Check each position
		for i, ch := range word {
			if row >= len(board) || col+i >= len(board[0]) {
				return false
			}
			cell := board[row][col+i]
			if cell != " " && strings.ToLower(cell) != string(ch) {
				return false
			}
		}
	}
	return true
}
