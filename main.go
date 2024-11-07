package main

import (
	"fmt"

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
	board:= tetris.CreateBoard(tetriminoes)
	fmt.Println(board)
	cleantetris := tetris.ReplaceHashes(tetriminoes)
	trimtetris := tetris.TrimAlltetriminos(cleantetris)
	fmt.Println(trimtetris)

}
