package playingPieces

func GetPlayingPieceX() IPlayingPiece {
	return &BasePlayingPiece{
		symbol: SYMBOL_X,
	}
}
