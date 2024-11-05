package tetris

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

func ReturnIndexNonDots(s string, index int) string {
	return s[:index] + s[index+1:]
}

func TrimVertical(tetrimino []string) []string {
	for j := len(tetrimino[0]) - 1; j >= 0; j-- {
		var notdot bool
		for i := range tetrimino {
			if tetrimino[i][j] != '.' {
				notdot = true
				break

			}
			if !notdot {
				tetrimino[i] = ReturnIndexNonDots(tetrimino[i], j)
			}
		}
	}
	return tetrimino
}
