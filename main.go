package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func getFileLineLength(fileName string) int {
	lineCount := 0
	file, _ := os.Open(fileName)
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lineCount++
	}
	defer file.Close()

	return lineCount
}

func getWordOnLineN(lineNumber int, fileName string) string {
	lastLine := 0
	file, _ := os.Open(fileName)
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lastLine++
		if lastLine == lineNumber {
			return fileScanner.Text()
		}
	}
	defer file.Close()

	return "Error"
}

func getLenOfLongestWord(words []string) int {
	longestWord := words[0]
	for i := range words {
		if len(longestWord) < len(words[i]) {
			longestWord = words[i]
		}
	}
	return len(longestWord)
}

func getLongestWord(words []string) []string {
	longestWord := words[0]
	for i := range words {
		if len(longestWord) < len(words[i]) {
			longestWord = words[i]
		}
	}
	wordSplittedToChars := strings.Split(longestWord, "")
	return wordSplittedToChars
}

func fillBoardWithWords(words []string) [][]string {
	leveys := getLenOfLongestWord(words)
	korkeus := getLenOfLongestWord(words)

	board := makeBoard(leveys, korkeus)

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
			}
			attempts--
		}
		if !placed {
			fmt.Println("Ei voitu laittaa sanaa:", word)
		}
	}

	fillBoardBlankSpaces(board)

	fmt.Println("\nBoard:")
	for _, row := range board {
		for _, ch := range row {
			fmt.Printf("%s ", ch)
		}
		fmt.Println()
	}

	return board
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

func fillBoardBlankSpaces(board [][]string) {
	for i := range board {
		for y := range board[i] {
			if board[i][y] == " " {
				board[i][y] = "*"
			}
		}
	}
}

func main() {
	fileName := "words.txt"

	lineCount := getFileLineLength(fileName)
	var randomLines []int
	wordCount := 15
	for i := 0; i < wordCount; i++ {
		randomLine := rand.Intn(lineCount)
		randomLines = append(randomLines, randomLine)
		i++
	}

	var randWords []string
	for i := range len(randomLines) {
		randWord := getWordOnLineN(randomLines[i], fileName)
		randWords = append(randWords, randWord)
		i++
	}
	fmt.Println(randWords)

	fillBoardWithWords(randWords)
}
