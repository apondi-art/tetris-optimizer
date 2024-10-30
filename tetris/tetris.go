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
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			if len(tetrimino) > 0 {
				tetriminoes = append(tetriminoes, tetrimino)
				tetrimino = nil
			}
			continue
		}

		tetrimino = append(tetrimino, input)
	}
	if len(tetrimino) > 0 {
		tetriminoes = append(tetriminoes, tetrimino)
	}
	return tetriminoes
}

func ReplaceHash(tetriminoes [][]string) [][]string {
	for i, tetrimino := range tetriminoes {
		for j, mino := range tetrimino {
			minoRunes := []rune(mino)
			for k := range minoRunes {
				if minoRunes[k] == '#' {
					minoRunes[k] = rune('A' + i)
				}
			}
			tetrimino[j] = string(minoRunes)
		}
	}
	return tetriminoes
}
