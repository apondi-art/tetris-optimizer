package tetris

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Readfile() [][]string {
	var tetrimino []string
	var tetriminoes [][]string
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("error opening file ,%s", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()
		tetrimino = append(tetrimino, strings.TrimSpace(input))
		if input == "" {
			tetriminoes = append(tetriminoes, tetrimino)
			tetrimino = nil
		}

	}
	if tetrimino != nil {
		tetriminoes = append(tetriminoes, tetrimino)
	}
	return tetriminoes
}
