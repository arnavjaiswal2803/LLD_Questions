import { PlayingPiece } from "./playingPieces/playingPiece.ts";

export class Player {
  private _name: string;
  private _playingPiece: PlayingPiece;

  constructor(name: string, piece: PlayingPiece) {
    this._name = name;
    this._playingPiece = piece;
  }

  get name() {
    return this._name;
  }

  get piece() {
    return this._playingPiece;
  }
}
