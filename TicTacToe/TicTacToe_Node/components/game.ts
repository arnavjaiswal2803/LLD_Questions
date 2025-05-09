import prompt from "prompt-sync";

import { Board } from "./board.ts";
import { Player } from "./player.ts";

export class Game {
  private _players: Player[];
  private _board: Board;

  constructor(size: number, players: Player[]) {
    this._board = new Board(size);
    this._players = players;
  }

  startGame() {
    const promptSync = prompt({ sigint: true });
    this._board.displayBoard();

    let turn = 0;
    while (this._board.areCellsAvailable()) {
      const player = this._players[turn];

      console.log(`----- ${player.name}'s turn -----`);
      const row = parseInt(promptSync("Enter row: "));
      const col = parseInt(promptSync("Enter col: "));

      if (!this._board.makeMove(row, col, player.piece)) {
        console.log("illegal move!!! Try again");
        continue;
      }

      this._board.displayBoard();

      if (this._board.hasWon()) {
        console.log(`${player.name} has won!!!`);
        return;
      }

      turn = (turn + 1) % this._players.length;
    }

    console.log("game tied");
  }
}
