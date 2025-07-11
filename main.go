package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
)

func getFileLineLength(filename string) int {
	lineCount := 0
	file, _ := os.Open(filename)
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lineCount++
	}
	defer file.Close()

	return lineCount
}

func getWordOnLineN(lineNumber int, filename string) string {
	lastLine := 0
	file, _ := os.Open(filename)
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
				fmt.Printf("Word %v starts in %v %v\n", word, row, col)
			}
			attempts--
		}
		if !placed {
			fmt.Println("Ei voitu laittaa sanaa:", word)
		}
	}

	fillBoardBlankSpaces(board)
	printBoard(board)
	return board
}

func printBoard(board [][]string) {
	fmt.Println("\nBoard:")
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
func getRandomChar() rune {
	charsList := []rune{}
	for rune := 'a'; rune <= 'z'; rune++ {
		charsList = append(charsList, rune)
	}
	charsList = append(charsList, 'ä', 'ö')
	randInt := rand.Intn(len(charsList))
	return charsList[randInt]
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

func removeItemFromList(list []string, item string) []string {
	for i, other := range list {
		if other == item {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

// func replaceFoundWordWithStars(word string) {
// 	wordPos :=
// }

func getRandomLines(filename string) []int {
	lineCount := getFileLineLength(filename)
	randomLines := []int{}
	wordCount := 15
	for i := 0; i < wordCount; i++ {
		randomLine := rand.Intn(lineCount)
		randomLines = append(randomLines, randomLine)
		i++
	}
	return randomLines
}

func getRandomWords(filename string) []string {
	randWords := []string{}
	randomLines := getRandomLines(filename)
	for i := range len(randomLines) {
		randWord := getWordOnLineN(randomLines[i], filename)
		randWords = append(randWords, randWord)
		i++
	}
	return randWords
}

func main() {
	filename := "../wordfind/words.txt"

	randWords := getRandomWords(filename)

	fillBoardWithWords(randWords)

	//Make copy of the word list I DONT KNOW WHY TODO: RESEARCH
	wordListCopy := make([]string, len(randWords))
	copy(wordListCopy, randWords)

	input := ""
	for len(wordListCopy) != 0 {
		fmt.Print("What words are in the puzzle >")
		fmt.Scan(&input)
		if slices.Contains(randWords, input) {
			fmt.Println("You found a word in the puzzle")
			wordListCopy = removeItemFromList(wordListCopy, input)
			// TODO: Replace right find with stars at the start with now
			// replaceFoundWordWithStars(input)
		} else {
			fmt.Println("Try again")
		}
	}
}
