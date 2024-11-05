package tetris

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadTetriminos reads tetrimino shapes from a file and returns them as a slice of slice of strings.
// ReadTetrimins read data from text file then append each block to the 2-D slice
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
