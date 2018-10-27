package game

import ."../board"
import ."../player"

const WINS = " WINS"
const TIE_GAME = "TIE GAME"
const NO_WINNER_YET = "NO WINNER YET"

func Status(board Board, currentPlayer Player) string {
   if board.WinningCombination() {
      return currentPlayer.Symbol + " (" + currentPlayer.Type + ")" + WINS
   } else if board.TieCombination() {
      return TIE_GAME
   }  
   return NO_WINNER_YET 
}
