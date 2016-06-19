package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println("Welcome to tictactoe \"_\" means the position is empty select from one of the position number belows to place your mark.")
	fmt.Println(1, 2, 3)
	fmt.Println(4, 5, 6)
	fmt.Println(7, 8, 9)
	fmt.Println("Player1 = \"X\"")
	fmt.Println("Player2 = \"O\"")
	playGame(board)
	fmt.Println("Game Over")
}

func playGame(board [][]string) [][]string {
	player1 := "X" 
	//player2 := "O"
	pos := selectLocation()
	invalidPosition(pos)
	turn(pos, board, player1)
	win := winner(board)
	if win == true {
		return board
	}
	pos = selectLocation()
	invalidPosition(pos)
	turn(pos, board, player2)
	win = winner(board)
	if win == true {
		return board
	}
	return playGame(board)
}
func invalidPosition(pos int) {
	if pos > 9 || pos < 1 {
		pos = tryAgain(pos)
	}
}
func tryAgain(pos int) int {
	if pos > 9 || pos < 1 {
		fmt.Println("please select a number between 1-9")
		return tryAgain(selectLocation())
	}
	return pos
}
func currentPosition(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func selectLocation() int {
	fmt.Println("Select a location on the board.")
	var pos int
	fmt.Scan(&pos)
	fmt.Println("Position ", pos, " Selected")
	return pos
}

func turn(pos int, board [][]string, player string) {
	if pos == (1) {
		board[0][0] = player
	}
	if pos == (2) {
		board[0][1] = player
	}
	if pos == (3) {
		board[0][2] = player
	}
	if pos == (4) {
		board[1][0] = player
	}
	if pos == (5) {
		board[1][1] = player
	}
	if pos == (6) {
		board[1][2] = player
	}
	if pos == (7) {
		board[2][0] = player
	}
	if pos == (8) {
		board[2][1] = player
	}
	if pos == (9) {
		board[2][2] = player
	}
	clear()
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func winner(board [][]string) bool {
	xWinsrow := []string{"X", "X", "X"}
	oWinsrow := []string{"O", "O", "O"}
	win := false
	switch {
	case reflect.DeepEqual(board[0], xWinsrow):
		fmt.Println("Player 1 wins !!")
		win = true
		return win
	case reflect.DeepEqual(board[1], xWinsrow):
		fmt.Println("Player 1 wins !!")
		win = true
		return win
	case reflect.DeepEqual(board[2], xWinsrow):
		fmt.Println("Player 1 wins !!")
		win = true
		return win
	case reflect.DeepEqual(board[0], oWinsrow):
		fmt.Println("Player 2 wins !!")
		win = true
		return win
	case reflect.DeepEqual(board[1], oWinsrow):
		fmt.Println("Player 2 wins !!")
		win = true
		return win
	case reflect.DeepEqual(board[2], oWinsrow):
		fmt.Println("Player 2 wins !!")
		win = true
		return win
	case board[0][0] == "X" && board[1][0] == "X" && board[2][0] == "X":
		win = true
		fmt.Println("Player 1 wins !")
		return win
	case board[0][1] == "X" && board[1][1] == "X" && board[2][1] == "X":
		win = true
		fmt.Println("Player 1 wins !")
		return win
	case board[0][2] == "X" && board[1][2] == "X" && board[2][2] == "X":
		win = true
		fmt.Println("Player 1 wins !")
		return win
	case board[0][0] == "O" && board[1][0] == "O" && board[2][0] == "O":
		win = true
		fmt.Println("Player 2 wins !")
		return win
	case board[0][1] == "O" && board[1][1] == "O" && board[2][1] == "O":
		win = true
		fmt.Println("Player 2 wins !")
		return win
	case board[0][2] == "O" && board[1][2] == "O" && board[2][2] == "O":
		win = true
		fmt.Println("Player 2 wins !")
		return win
	case board[0][0] == "X" && board[1][1] == "X" && board[2][2] == "X":
		win = true
		fmt.Println("Player 1 wins !")
	case board[0][2] == "X" && board[1][1] == "X" && board[2][7] == "X":
		win = true
		fmt.Println("Player 1 wins !")
	case board[0][0] == "O" && board[1][1] == "O" && board[2][2] == "O":
		win = true
		fmt.Println("Player 2 wins !")
	case board[0][2] == "O" && board[1][1] == "O" && board[2][7] == "O":
		win = true
		fmt.Println("Player 2 wins !")
	}
	return win
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
