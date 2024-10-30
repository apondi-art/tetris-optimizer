package main

import (
	"fmt"

	"learn.zone01kisumu.ke/git/quochieng/tetris-optimizer.git/tetris"
)

func main() {
	result := tetris.Readfile()
	final := tetris.ReplaceHash(result)
	fmt.Println(final)
}
