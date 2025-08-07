package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func fillBoardWithWords(board [][]string, words []string, leveys, korkeus int) ([][]string, []string) {
	remainingWords := make([]string, len(words))
	copy(remainingWords, words)
	placedWords := []string{}

	for i := 0; i < len(remainingWords); i++ {
		word := strings.ToLower(remainingWords[i]) // Ensure consistent case
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
				placedWords = append(placedWords, word)
				fmt.Printf("Word %v starts in %v %v\n", word, row, col)
			}
			attempts--
		}
		if !placed {
			fmt.Println("Ei voitu laittaa sanaa:", word)
			remainingWords = append(remainingWords[:i], remainingWords[i+1:]...)
			i--
		}
	}
	return board, placedWords
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

func replaceFoundWordWithStars(board [][]string, word string) bool {
	word = strings.ToLower(word)
	wordLen := len(word)
	found := false

	for row := range board {
		for col := 0; col <= len(board[row])-wordLen; col++ {
			match := true
			for i := range wordLen {
				if strings.ToLower(board[row][col+i]) != string(word[i]) {
					match = false
					break
				}
			}
			if match {
				for i := range wordLen {
					board[row][col+i] = "*"
				}
				found = true
				return found
			}
		}
	}

	for col := 0; col < len(board[0]); col++ {
		for row := 0; row <= len(board)-wordLen; row++ {
			match := true
			for i := range wordLen {
				if strings.ToLower(board[row+i][col]) != string(word[i]) {
					match = false
					break
				}
			}
			if match {
				for i := range wordLen {
					board[row+i][col] = "*"
				}
				found = true
				return found
			}
		}
	}
	return found
}
