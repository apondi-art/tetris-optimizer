package tetris


func Place(board [][]string, tetrimino []string, x, y int) {
	for dy := range tetrimino {
		for dx, char := range tetrimino[dy] {
			board[dy+y][dx+x] = string(char)

		}
	}

}

func CanPlace(board [][]string, tetrimino []string, x, y int) bool {
    for dy := range tetrimino {
        for dx := range tetrimino[dy] {
            // Calculate the position on the board
            boardY, boardX := dy+y, dx+x

            // Check if the position is outside the bounds of the board
            if boardY >= len(board) || boardX >= len(board[0]) || boardY < 0 || boardX < 0 {
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


func RemoveTetris(board [][]string, tetrimino []string, x,y int){
	for dy := range tetrimino{
		for dx :=range tetrimino[dy]{
			board[dy][dx] = "."
		}
	}
}

