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
	fmt.Println("number of lines:", lineCount)
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

func main() {
	//TODO: vaihda myöhemmin johonkin
	var korkeus int
	var leveys int

	fmt.Println("Reading file words")
	fileName := "words.txt"
	/**content**/ _, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	lineCount := getFileLineLength(fileName)
	var randomLines []int
	wordCount := 10
	for i := range wordCount {
		randomLine := rand.Intn(lineCount)
		randomLines = append(randomLines, randomLine)
		i++
	}
	fmt.Println(randomLines)

	var randWords []string
	for i := range len(randomLines) {
		randWord := getWordOnLineN(randomLines[i], fileName)
		randWords = append(randWords, randWord)
		i++
	}
	fmt.Println(randWords)

	leveys = getLenOfLongestWord(randWords)
	korkeus = getLenOfLongestWord(randWords)

	board := make([][]string, korkeus)
	for i := range board {
		board[i] = make([]string, leveys)
		for j := range board[i] {
			board[i][j] = "[]"
		}
	}

	// Put biggest word on board
	// Get random point on y-axis
	// start word on that
	randomRow := rand.Intn(korkeus)
	longestWord := getLongestWord(randWords)
	for i, char := range longestWord {
		board[randomRow][i] = string(char)
	}

	fmt.Println("\nBoard:")
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
}
