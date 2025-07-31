package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"strings"

	templates "github.com/NuuttiSir/wordfind/internal/templates"
	_ "github.com/a-h/templ"
)

var (
	allWords     []string // All words that were placed on the board
	currentWords []string // Words still to be found
	leveys       int
	korkeus      int
	board        [][]string
)

func main() {

	const port = "6969"

	initializeGame()

	//TODO: make different languages

	http.HandleFunc("/", handler_homePage)
	http.HandleFunc("/submit", handler_Submit)
	http.HandleFunc("/reset", handler_resetBoard)

	fmt.Println("Listening on 6969")
	http.ListenAndServe(":6969", nil)
}

func initializeGame() {
	words := getRandomWords("./words.txt")
	leveys = getLenOfLongestWord(words) + 2
	korkeus = getLenOfLongestWord(words) + 2

	// Ensure minimum board size
	if leveys < 10 {
		leveys = 10
	}
	if korkeus < 10 {
		korkeus = 10
	}

	board = makeBoard(leveys, korkeus)
	var placedWords []string
	board, placedWords = fillBoardWithWords(board, words, leveys, korkeus)
	fillBoardBlankSpaces(board)

	// Only track words that were actually placed
	allWords = make([]string, len(placedWords))
	copy(allWords, placedWords)
	currentWords = make([]string, len(placedWords))
	copy(currentWords, placedWords)

	fmt.Printf("Game initialized with %d words\n", len(currentWords))
}

func handler_homePage(w http.ResponseWriter, r *http.Request) {
	templates.MainPage(board).Render(r.Context(), w)
}

func handler_Submit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	word := strings.ToLower(strings.TrimSpace(r.FormValue("word")))

	if word == "" {
		templates.PlayPage(board).Render(r.Context(), w)
		return
	}

	// Check if word exists in current words (words still to be found)
	if slices.Contains(currentWords, word) {
		currentWords = removeItemFromList(currentWords, word)
		wordFound := replaceFoundWordWithStars(board, word)

		if wordFound {
			fmt.Printf("Word found: %s\n", word)
		} else {
			fmt.Printf("Word %s was in list but not found on board\n", word)
		}
	} else {
		fmt.Printf("Word not found or already guessed: %s\n", word)
	}

	// Check win condition
	if len(currentWords) == 0 {
		templates.WinPage().Render(r.Context(), w)
		return
	}
	templates.PlayPage(board).Render(r.Context(), w)
}

func handler_resetBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Board reset requested")
	initializeGame()
	templates.PlayPage(board).Render(r.Context(), w)
}

func getFileLineLength(wordsFile string) int {
	lineCount := 0
	file, _ := os.Open(wordsFile)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lineCount++
	}

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
