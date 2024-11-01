package main

import (
	"fmt"
	"strings"

	"learn.zone01kisumu.ke/git/quochieng/tetris-optimizer.git/tetris"
)

func main() {
	// result := tetris.Readfile()
	// final := tetris.ReplaceHash(result)
	// fmt.Println(final)
	s := "...A ...A ...A ...A"
	str := (strings.Fields(s))
	fmt.Println(tetris.CheckingValidTetrimino(str))
}
