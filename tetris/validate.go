package tetris

import "fmt"

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
