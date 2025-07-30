package main

import "math/rand"

func getRandomChar() rune {
	charsList := []rune{}
	for rune := 'a'; rune <= 'z'; rune++ {
		charsList = append(charsList, rune)
	}
	charsList = append(charsList, 'ä', 'ö')
	randInt := rand.Intn(len(charsList))
	return charsList[randInt]
}

func getRandomLines(wordsFile string) []int {
	lineCount := getFileLineLength(wordsFile)
	randomLines := []int{}
	wordCount := 15
	for range wordCount {
		randomLine := rand.Intn(lineCount)
		randomLines = append(randomLines, randomLine)
	}
	return randomLines
}

func getRandomWords(wordsFile string) []string {
	randWords := []string{}
	randomLines := getRandomLines(wordsFile)
	for i := range randomLines {
		randWord := getWordOnLineN(randomLines[i], wordsFile)
		randWords = append(randWords, randWord)
	}
	return randWords
}

func getRandomLine(wordsFile string) int {
	lineCount := getFileLineLength(wordsFile)
	return rand.Intn(lineCount)
}

func getRandomWord(wordsFile string) string {
	return getWordOnLineN(getRandomLine(wordsFile), wordsFile)
}
