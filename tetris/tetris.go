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

func CheckingValidTetrimino(tetrimino []string) bool {
	countConnection := 0
	count := 0

	// Helper function to check if a given position is valid and not '.'
	isValid := func(i, j int) bool {
		return i >= 0 && i < len(tetrimino) && j >= 0 && j < len(tetrimino[i]) && tetrimino[i][j] != '.'
	}
	for i := range tetrimino {
		for j := range tetrimino[i] {
			if (tetrimino[i][j] >= 'A' && tetrimino[i][j] <= 'Z') {
				count++
			}
			if tetrimino[i][j] != '.' {
				// Check all four possible adjacent positions
				if isValid(i+1, j) { // Below
					countConnection++
				}
				if isValid(i-1, j) { // Above
					countConnection++
				}
				if isValid(i, j+1) { // Right
					countConnection++
				}
				if isValid(i, j-1) { // Left
					countConnection++
				}
			}
		}
	}
	if count == 4 && (countConnection >= 6 && countConnection <= 8) {
		return true

	}

	return false
}
