package main

import (
	"fmt"
	"math/rand"
)

func fillBoardWithWords(board [][]string, words []string, leveys, korkeus int) [][]string {
	for _, word := range words {
		placed := false
		attempts := 100
		for attempts > 0 && !placed {
			vertical := rand.Intn(2) == 0

			var row, col int
			if vertical {
				row = rand.Intn(korkeus - len(word) + 1)
				col = rand.Intn(leveys)
			} else {
				row = rand.Intn(korkeus)
				col = rand.Intn(leveys - len(word) + 1)
			}

			if canPlaceWord(board, word, row, col, vertical) {
				placeWord(board, word, row, col, vertical)
				placed = true
				fmt.Printf("Word %v starts in %v %v\n", word, row, col)
			}
			attempts--
		}
		if !placed {
			// TODO: make so it tries a new word
			// NOTE: ATM it just repeats the last word
			// Maybe add arandom word?
			removeItemFromList(words, word)
			fmt.Println("Removed word: ", word)
			fmt.Println(words)
			fmt.Println("Ei voitu laittaa sanaa:", word)
		}
	}
	return board
}

func printBoard(board [][]string) {
	for _, row := range board {
		for _, ch := range row {
			fmt.Printf("[%s]", ch)
		}
		fmt.Println()
	}
}

func makeBoard(leveys, korkeus int) [][]string {
	board := make([][]string, korkeus)
	for i := range board {
		board[i] = make([]string, leveys)
		for j := range board[i] {
			board[i][j] = " "
		}
	}
	return board
}

func fillBoardBlankSpaces(board [][]string) {
	for i := range board {
		for y := range board[i] {
			if board[i][y] == " " {
				board[i][y] = string(getRandomChar())
			}
		}
	}
}

func replaceFoundWordWithStars(board [][]string, word string) {
	wordLen := len(word)

	for row := range board {
		for col := 0; col <= len(board[row])-wordLen; col++ {
			match := true
			for i := range wordLen {
				if board[row][col+i] != string(word[i]) {
					match = false
					break
				}
			}
			if match {
				for i := range wordLen {
					board[row][col+i] = "*"
				}
				return
			}
		}
	}

	for col := 0; col < len(board[0]); col++ {
		for row := 0; row <= len(board)-wordLen; row++ {
			match := true
			for i := range wordLen {
				if board[row+i][col] != string(word[i]) {
					match = false
					break
				}
			}
			if match {
				for i := range wordLen {
					board[row+i][col] = "*"
				}
				return
			}
		}
	}
}
