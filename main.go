package main

import (
	"bufio"
	"os"
	"os/exec"

	_ "github.com/a-h/templ"
)

func main() {

	const port = "6969"

	wordsFile := "../wordfind/words.txt"
	//TODO: make different languages
	Play(wordsFile)
}

func getFileLineLength(wordsFile string) int {
	lineCount := 0
	file, _ := os.Open(wordsFile)
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lineCount++
	}
	defer file.Close()

	return lineCount
}

func removeItemFromList(list []string, item string) []string {
	for i, other := range list {
		if other == item {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
