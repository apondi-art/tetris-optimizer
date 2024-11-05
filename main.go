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

	tetriminoes = tetris.ReplaceHashes(tetriminoes)
	fmt.Println("Validated and transformed tetriminoes:", tetriminoes)
	for i := range tetriminoes {
		fmt.Println(tetris.TrimDotsHorizontal(tetriminoes[i]))
		// fmt.Println(tetris.TrimDotsHorizontal(tetriminoes[i]))
	}
}
