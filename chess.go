package main

import "fmt"

type chessboard [8][8]string

func main() {
	var board chessboard

	completeBoard := setupBlack(setupWhite(board))

	completeBoard.printBoard()
}

func setupWhite(board chessboard) chessboard {
	board[0][0] = "white rook"
	board[0][1] = "white knight"
	board[0][2] = "white bishop"
	board[0][3] = "white king"
	board[0][4] = "white queen"
	board[0][5] = "white bishop"
	board[0][6] = "white knight"
	board[0][7] = "white rook"

	for column := range board[1] {
		board[1][column] = "white pawn"
	}

	return board
}

func setupBlack(board chessboard) chessboard {
	board[7][0] = "black rook"
	board[7][1] = "black knight"
	board[7][2] = "black bishop"
	board[7][3] = "black king"
	board[7][4] = "black queen"
	board[7][5] = "black bishop"
	board[7][6] = "black knight"
	board[7][7] = "black rook"

	for column := range board[1] {
		board[6][column] = "black pawn"
	}

	return board
}

func (board chessboard) printBoard() {
	const line string = " --------------------------------------------------------------------------------------------------------------------------------------"

	fmt.Println(line)

	for _, row := range board {
		for _, column := range row {
			if column == "" {
				fmt.Printf("| %-14v|", " ")
			} else {
				fmt.Printf("| %-14v|", column)
			}
		}
		fmt.Println()
		fmt.Println(line)
	}

}
