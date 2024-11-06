package main

import (
	"fmt"
	"math"

	"learn.zone01kisumu.ke/git/quochieng/tetris-optimizer.git/tetris"
)

func main() {
	tetriminoes, err := tetris.ReadTetriminos("text.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !tetris.ValidateTetriminos(tetriminoes) {

		fmt.Println("Some tetrimino blocks are invalid.")
		return
	}

	tetriminoes = tetris.ReplaceHashes(tetriminoes)
	fmt.Println("Validated and transformed tetriminoes:", tetriminoes)
	for i := range tetriminoes {
		tetriminoes[i] = tetris.TrimDotsHorizontal(tetriminoes[i])
		tetriminoes[i] = tetris.TrimVertical(tetriminoes[i])
		fmt.Println(tetriminoes[i])
		// fmt.Println(tetris.TrimDotsHorizontal(tetriminoes[i]))
	}
}

func CreateBoard(tetriminos [][]string) [][]string {

	numTetriminos := len(tetriminos)

	size := int(math.Ceil(math.Sqrt(float64(numTetriminos * 4))))

	// Creating the board with the corrected size
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
	}

	// Initializing the board with "." as an empty cell
	for i := range board {
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}
