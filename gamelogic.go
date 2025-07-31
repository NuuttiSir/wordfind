package main

//
// import (
// 	"fmt"
// 	"slices"
// )
//
// func Play(filename string) {
// 	randWords := getRandomWords(filename)
// 	leveys := getLenOfLongestWord(randWords)
// 	korkeus := getLenOfLongestWord(randWords)
//
// 	board := makeBoard(leveys, korkeus)
// 	board = fillBoardWithWords(board, randWords, leveys, korkeus)
// 	fillBoardBlankSpaces(board)
// 	printBoard(board)
//
// 	//Make copy of the word list I DONT KNOW WHY TODO: RESEARCH
// 	wordListCopy := make([]string, len(randWords))
// 	copy(wordListCopy, randWords)
//
// 	input := ""
// 	for len(wordListCopy) != 0 {
// 		fmt.Print("What words are in the puzzle >")
// 		fmt.Scan(&input)
// 		if slices.Contains(randWords, input) {
// 			wordListCopy = removeItemFromList(wordListCopy, input)
// 			replaceFoundWordWithStars(board, input)
// 			clearTerminal()
// 			fmt.Println("You found a word in the puzzle")
// 			printBoard(board)
// 		} else {
// 			fmt.Println("Try again")
// 		}
// 	}
// }
