package components

import "ticTacToe/components/playingPieces"

func GetNewPlayer(name string, piece playingPieces.IPlayingPiece) *Player {
	return &Player{
		name:         name,
		playingPiece: piece,
	}
}

type Player struct {
	name         string
	playingPiece playingPieces.IPlayingPiece
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPlayingPiece() playingPieces.IPlayingPiece {
	return p.playingPiece
}
