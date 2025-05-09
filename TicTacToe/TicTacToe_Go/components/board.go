package components

import (
	"fmt"
	"ticTacToe/components/playingPieces"
)

func GetNewBoard(size int) *Board {
	board := make([][]playingPieces.IPlayingPiece, size)
	for i := 0; i < size; i++ {
		board[i] = make([]playingPieces.IPlayingPiece, size)
	}

	return &Board{
		board:         board,
		size:          size,
		occupiedCount: 0,
	}
}

type Board struct {
	board         [][]playingPieces.IPlayingPiece
	size          int
	occupiedCount int
}

func (b *Board) MakeMove(row, col int, piece playingPieces.IPlayingPiece) bool {
	if row < 0 || row >= b.size || col < 0 || col >= b.size ||
		!b.AreCellsAvailable() || b.board[row][col] != nil {

		return false
	}
	b.board[row][col] = piece
	b.occupiedCount++
	return true
}

func (b *Board) AreCellsAvailable() bool {
	return b.occupiedCount < (b.size * b.size)
}

func (b *Board) HasWon() bool {
	// check horizontally
	for i := 0; i < b.size; i++ {
		for j := 1; j < b.size-1; j++ {
			if b.board[i][j-1] != nil && b.board[i][j] != nil && b.board[i][j+1] != nil &&
				b.board[i][j-1].GetSymbol() == b.board[i][j].GetSymbol() &&
				b.board[i][j].GetSymbol() == b.board[i][j+1].GetSymbol() {
				return true
			}
		}
	}

	// check vertically
	for i := 0; i < b.size; i++ {
		for j := 1; j < b.size-1; j++ {
			if b.board[j-1][i] != nil && b.board[j][i] != nil && b.board[j+1][i] != nil &&
				b.board[j-1][i].GetSymbol() == b.board[j][i].GetSymbol() &&
				b.board[j][i].GetSymbol() == b.board[j+1][i].GetSymbol() {
				return true
			}
		}
	}

	// check diagonal
	for i := 1; i < b.size-1; i++ {
		if b.board[i-1][i-1] != nil && b.board[i][i] != nil && b.board[i+1][i+1] != nil &&
			b.board[i-1][i-1].GetSymbol() == b.board[i][i].GetSymbol() &&
			b.board[i][i].GetSymbol() == b.board[i+1][i+1].GetSymbol() {

			return true
		}
	}

	// check anti-diagonal
	for i := 1; i < b.size-1; i++ {
		if b.board[b.size-(i-1)-1][i-1] != nil && b.board[b.size-i-1][i] != nil && b.board[b.size-(i+1)-1][i+1] != nil &&
			b.board[b.size-(i-1)-1][i-1].GetSymbol() == b.board[b.size-i-1][i].GetSymbol() &&
			b.board[b.size-i-1][i].GetSymbol() == b.board[b.size-(i+1)-1][i+1].GetSymbol() {

			return true
		}
	}
	return false
}

func (b *Board) DisplayBoard() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.board[i][j] == nil {
				fmt.Print("nil")
			} else {
				fmt.Print(b.board[i][j].GetSymbol())
			}

			if j < b.size-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
	}
}
