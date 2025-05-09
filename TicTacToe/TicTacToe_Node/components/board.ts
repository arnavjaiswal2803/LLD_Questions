import { PlayingPiece } from "./playingPieces/playingPiece.ts";

export class Board {
  private _board: PlayingPiece[][];
  private _size: number;
  private _occupiedCount: number;

  constructor(size: number) {
    this._size = size;
    this._occupiedCount = 0;
    this._board = Array.from({ length: size }, () => Array(size).fill(null));
  }

  areCellsAvailable() {
    return this._occupiedCount < this._size * this._size;
  }

  makeMove(row: number, col: number, piece: PlayingPiece) {
    if (
      row < 0 ||
      row >= this._size ||
      col < 0 ||
      col >= this._size ||
      this._board[row][col] != null
    ) {
      return false;
    }

    this._board[row][col] = piece;
    this._occupiedCount++;
    return true;
  }

  hasWon() {
    // check horizontally
    for (let row = 0; row < this._size; row++) {
      for (let col = 1; col < this._size - 1; col++) {
        if (
          this._board[row][col - 1] !== null &&
          this._board[row][col] !== null &&
          this._board[row][col + 1] !== null &&
          this._board[row][col - 1].symbol === this._board[row][col].symbol &&
          this._board[row][col].symbol === this._board[row][col + 1].symbol
        ) {
          return true;
        }
      }
    }

    // check vertically
    for (let col = 0; col < this._size; col++) {
      for (let row = 1; row < this._size - 1; row++) {
        if (
          this._board[row - 1][col] !== null &&
          this._board[row][col] !== null &&
          this._board[row + 1][col] !== null &&
          this._board[row - 1][col].symbol === this._board[row][col].symbol &&
          this._board[row][col].symbol === this._board[row + 1][col].symbol
        ) {
          return true;
        }
      }
    }

    // check diagonally
    for (let i = 1; i < this._size - 1; i++) {
      if (
        this._board[i - 1][i - 1] !== null &&
        this._board[i][i] !== null &&
        this._board[i + 1][i + 1] !== null &&
        this._board[i - 1][i - 1].symbol === this._board[i][i].symbol &&
        this._board[i][i].symbol === this._board[i + 1][i + 1].symbol
      ) {
        return true;
      }
    }

    // check anti-diagonally
    for (let i = 1; i < this._size - 1; i++) {
      const n = this._size - 1;
      if (
        this._board[n - (i - 1)][i - 1] !== null &&
        this._board[n - i][i] !== null &&
        this._board[n - (i + 1)][i + 1] !== null &&
        this._board[n - (i - 1)][i - 1].symbol ===
          this._board[n - i][i].symbol &&
        this._board[n - i][i].symbol === this._board[n - (i + 1)][i + 1].symbol
      ) {
        return true;
      }
    }

    return false;
  }

  displayBoard() {
    for (let i = 0; i < this._size; i++) {
      for (let j = 0; j < this._size; j++) {
        if (this._board[i][j] === null) {
          process.stdout.write("null");
        } else {
          process.stdout.write(`${this._board[i][j].symbol}`);
        }
        if (j < this._size - 1) process.stdout.write(` | `);
      }
      console.log("");
    }
  }
}
