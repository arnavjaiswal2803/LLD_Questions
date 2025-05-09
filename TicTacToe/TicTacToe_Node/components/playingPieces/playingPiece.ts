export const Symbol = {
  X: "X",
  O: "O",
};

export type Symbol = (typeof Symbol)[keyof typeof Symbol];

export abstract class PlayingPiece {
  private _piece: Symbol;

  constructor(piece: Symbol) {
    this._piece = piece;
  }

  get symbol() {
    return this._piece;
  }
}
