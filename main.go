package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"slices"

	templates "github.com/NuuttiSir/wordfind/internal/templates"
	_ "github.com/a-h/templ"
)

var words = getRandomWords("./words.txt")
var leveys = getLenOfLongestWord(words)
var korkeus = getLenOfLongestWord(words)
var board = makeBoard(leveys, korkeus)

type Board struct {
	board [][]string
}

func main() {

	const port = "6969"

	//wordsFile := "../wordfind/words.txt"
	//TODO: make different languages
	//Play(wordsFile)
	board = fillBoardWithWords(board, words, leveys, korkeus)
	fillBoardBlankSpaces(board)

	http.HandleFunc("/", handler_HomePage)
	http.HandleFunc("/submit", handler_Submit)
	fmt.Println("Listening on 6969")
	http.ListenAndServe(":6969", nil)
}

func handler_HomePage(w http.ResponseWriter, r *http.Request) {
	templates.MainPage(board).Render(r.Context(), w)
}

func handler_Submit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	word := r.FormValue("word")
	//for len(words) != 0 {
	if slices.Contains(words, word) {
		words = removeItemFromList(words, word)
		replaceFoundWordWithStars(board, word)
	} else {
		fmt.Println("Try again")
	}
	templates.BoardPrint(board).Render(r.Context(), w)
	//}
	// fmt.Println("VOITTO")
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
