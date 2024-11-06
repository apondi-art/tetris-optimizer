package tetris

// Function to check if a Tetrimino can be placed at position (x, y) on the board
func CanPlace(board [][]string, tetrimino []string, x, y int) bool {
	for dy := range tetrimino {
		for dx := range tetrimino[dy] {
			// Calculate the position on the board
			boardY, boardX := dy+y, dx+x

			// Check if the position is outside the bounds of the board
			if boardY < 0 || boardY >= len(board) || boardX < 0 || boardX >= len(board[0]) {
				return false
			}

			// Check if the board position is already occupied
			if board[boardY][boardX] != "." {
				return false
			}
		}
	}
	return true
}

// Function to place a Tetrimino on the board at position (x, y)
func Place(board [][]string, tetrimino []string, x, y int) {
	for dy := range tetrimino {
		for dx, char := range tetrimino[dy] {
			// Place the character on the board
			board[dy+y][dx+x] = string(char)
		}
	}
}

// Function to remove a Tetrimino from the board at position (x, y)
func RemoveTetris(board [][]string, tetrimino []string, x, y int) {
	for dy := range tetrimino {
		for dx := range tetrimino[dy] {
			// Calculate the position on the board
			boardY, boardX := dy+y, dx+x

			// Ensure the position is within bounds before attempting to remove
			if boardY >= 0 && boardY < len(board) && boardX >= 0 && boardX < len(board[0]) {
				board[boardY][boardX] = "."
			}
		}
	}
}

func PlaceRecursively(board [][]string, tetriminos [][]string, index int) bool {
	// Base case: If index is equal to the number of Tetriminos, all have been placed
	if index == len(tetriminos) {
		return true
	}

	tetrimino := tetriminos[index] // Current Tetrimino piece to be placed
	for y := range board {
		for x := range board[y] {
			if CanPlace(board, tetrimino, x, y) {
				Place(board, tetrimino, x, y) // Place current Tetrimino
				// Recurse with the next Tetrimino
				if PlaceRecursively(board, tetriminos, index+1) {
					return true // If successful, keep the placement
				}
				RemoveTetris(board, tetrimino, x, y) // Otherwise, backtrack
			}
		}
	}
	return false // If no valid placement found, return false
}
