package main

import (
	"bufio"
	"os"
)

func getWordOnLineN(lineNumber int, wordsFile string) string {
	lastLine := 0
	file, _ := os.Open(wordsFile)
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

// func getLongestWord(words []string) []string {
// 	longestWord := words[0]
// 	for i := range words {
// 		if len(longestWord) < len(words[i]) {
// 			longestWord = words[i]
// 		}
// 	}
// 	wordSplittedToChars := strings.Split(longestWord, "")
// 	return wordSplittedToChars
// }
