package tetris

import "math"

func BoardSize(tetris [][]string) int {
	size := int(math.Ceil(math.Sqrt(float64(len(tetris) * 4))))
	return size
}

func TetrisBoard(size int) [][]string {
	wholeBoard := make([][]string, size)
	for i := range wholeBoard {
		wholeBoard[i] = make([]string, size)
		for j := range wholeBoard[i] {
			wholeBoard[i][j] = "."
		}

	}
	return wholeBoard
}

// ReplaceHashes replaces '#' characters in tetrimino blocks with letters 'A', 'B', etc.
func ReplaceHashes(tetriminoes [][]string) [][]string {
	for i, tetrimino := range tetriminoes {
		for j, row := range tetrimino {
			rowRunes := []rune(row)
			for k, char := range rowRunes {
				if char == '#' {
					rowRunes[k] = rune('A' + i)
				}
			}
			tetrimino[j] = string(rowRunes)
		}
	}
	return tetriminoes
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
