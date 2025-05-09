import { Game } from "./components/game.ts";
import { Player } from "./components/player.ts";
import { PlayingPieceX } from "./components/playingPieces/playingPieceX.ts";
import { PlayingPieceO } from "./components/playingPieces/playingPieceO.ts";

function main() {
  const players: Player[] = [];
  players.push(new Player("Foo", new PlayingPieceX()));
  players.push(new Player("Bar", new PlayingPieceO()));

  const game = new Game(3, players);
  game.startGame();
}

main();
