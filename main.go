package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	templates "github.com/NuuttiSir/wordfind/internal/templates"
	_ "github.com/a-h/templ"
)

func main() {

	const port = "6969"

	//wordsFile := "../wordfind/words.txt"
	//TODO: make different languages
	//Play(wordsFile)

	http.HandleFunc("/", handler_HomePage)
	fmt.Println("Listening on 6969")
	http.ListenAndServe(":6969", nil)
}

func handler_HomePage(w http.ResponseWriter, r *http.Request) {
	words := getRandomWords("./words.txt")
	leveys := getLenOfLongestWord(words)
	korkeus := getLenOfLongestWord(words)

	board := makeBoard(leveys, korkeus)
	board = fillBoardWithWords(board, words, leveys, korkeus)
	fillBoardBlankSpaces(board)
	err := templates.MainPage(board).Render(r.Context(), w)
	if err != nil {
		fmt.Println("VITTU")
	}

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
