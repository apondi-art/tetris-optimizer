package tetris

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadTetriminos reads tetrimino shapes from a file and returns them as a slice of slice of strings.
//ReadTetrimins read data from text file then append each block to the 2-D slice
func ReadTetriminos(filename string) ([][]string, error) {
	var tetrimino []string
	var tetriminoes [][]string
	// Open the file and handle any potential errors
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check for an empty line, used to separate tetrimino blocks
		if line == "" {
			if len(tetrimino) > 0 {
				tetriminoes = append(tetriminoes, tetrimino)
				tetrimino = nil
			}
			continue
		}

		// Validate the line length and characters
		if !isValidLength(line) {
			return nil, fmt.Errorf("error: line length must be 4 characters")
		}
		if !hasValidCharacters(line) {
			return nil, fmt.Errorf("error: invalid characters, only '#' and '.' are allowed")
		}


		// Add valid line to the current tetrimino block
		tetrimino = append(tetrimino, line)
	}

	// Check for any remaining tetrimino block at the end of the file
	if len(tetrimino) > 0 {
		tetriminoes = append(tetriminoes, tetrimino)
	}

	return tetriminoes, nil
}

// isValidLength checks if the length of the string is equal to 4
func isValidLength(s string) bool {
	return len(s) == 4
}

// hasValidCharacters checks if the string contains only '#' or '.'
func hasValidCharacters(s string) bool {
	for _, char := range s {
		if char != '#' && char != '.' {
			return false
		}
	}
	return true
}

// ValidateTetriminos checks if all tetrimino blocks are valid
func ValidateTetriminos(tetriminoes [][]string) bool {
	for _, tetrimino := range tetriminoes {
		if !isTetriminoValid(tetrimino) {
			fmt.Println("Error: Invalid tetrimino")
			return false
		}
	}
	return true
}

// isTetriminoValid checks if a single tetrimino block is valid
func isTetriminoValid(tetrimino []string) bool {
	connectionCount := 0
	hashCount := 0

	// Helper function to check if a given position is valid and not '.'
	isConnected := func(i, j int) bool {
		return i >= 0 && i < len(tetrimino) && j >= 0 && j < len(tetrimino[i]) && tetrimino[i][j] == '#'
	}

	// Iterate over the tetrimino grid to count '#' characters and connections
	for i, row := range tetrimino {
		for j, char := range row {
			if char == '#' {
				hashCount++
				// Check all four possible adjacent positions for connections
				if isConnected(i+1, j) { // Below
					connectionCount++
				}
				if isConnected(i-1, j) { // Above
					connectionCount++
				}
				if isConnected(i, j+1) { // Right
					connectionCount++
				}
				if isConnected(i, j-1) { // Left
					connectionCount++
				}
			}
		}
	}

	// A valid tetrimino must have exactly 4 '#' characters and 6-8 connections
	return hashCount == 4 && connectionCount >= 6 && connectionCount <= 8
}

// func BoardSize(tetris [][]string) int {
// 	size := int(math.Ceil(math.Sqrt(float64(len(tetris) * 4))))
// 	return size
// }

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

func CheckDotsHorizontal(row string) bool {
	// Returns true if the row contains only dots
	for _, char := range row {
		if char != '.' {
			return false
		}
	}
	return true
}

func TrimDotsHorizontal(tetrimino []string) []string {
	var trimmed []string

	// Iterate through each row of the tetrimino
	for i := range tetrimino {
		// If the row does not contain only dots, add it to the trimmed slice
		if !CheckDotsHorizontal(tetrimino[i]) {
			trimmed = append(trimmed, tetrimino[i])
		}
	}

	// Return the rows without dots
	return trimmed
}

func ReturnIndexNonDots(s string, index int)string{
	return s[:index] + s[index+1:]
}


func TrimVertical(tetrimino []string)[]string{
	for j:= len(tetrimino[0])-1; j>= 0; j--{
		var notdot bool
		for i:=range tetrimino{
			if tetrimino[i][j] != '.'{
				notdot = true
				break

			}
			if !notdot{
				tetrimino[i] = ReturnIndexNonDots(tetrimino[i],j)
			
			} 
		}
	}
	return tetrimino

}


