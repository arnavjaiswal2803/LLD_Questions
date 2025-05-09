package playingPieces

type Symbol string

const (
	SYMBOL_X Symbol = "X"
	SYMBOL_O Symbol = "O"
)

type IPlayingPiece interface {
	GetSymbol() Symbol
}

type BasePlayingPiece struct {
	symbol Symbol
}

func (b *BasePlayingPiece) GetSymbol() Symbol {
	return b.symbol
}
