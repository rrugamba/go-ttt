package game

import "../board"
import "../player"

func Status(b board.Board, currentPlayer player.Player) string {
   if b.WinningCombination() {
      return currentPlayer.Symbol + " WINS"
   } else if b.TieCombination() {
      return "TIE GAME"
   }  
   return "NO WINNER YET" 
}
