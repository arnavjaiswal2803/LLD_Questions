package playingPieces

func GetPlayingPieceO() IPlayingPiece {
	return &BasePlayingPiece{
		symbol: SYMBOL_O,
	}
}
