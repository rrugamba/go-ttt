package moveStrategy

import "../board"
import "../player"
import "strconv"

func GetMove(a board.Board, d player.Player) int {
    return WhoseMove(a, d)
}

func WhoseMove(a board.Board, d player.Player) int {
      if d.Type == "human" {
         return HumanMove(a)
      } else {
         return AiMove(a, d)
      }
}

func HumanMove(a board.Board) int {
      for i := 0; i < len(a.Array); i++ {
          if a.Array[i] == strconv.Itoa(i+1) {
	     return i
          }
      }
      return -1
}

func AiMove(a board.Board, d player.Player) int {
     return 2	 
}
