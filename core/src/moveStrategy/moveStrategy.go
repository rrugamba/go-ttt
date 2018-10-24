package moveStrategy

import "../board"
import "../player"
import "strconv"

func GetMove(a board.Board, d player.Player, e player.Player) int {
    return WhoseMove(a, d, e)
}

func WhoseMove(a board.Board, d player.Player,  e player.Player) int {
      if d.Type == "human" {
         return HumanMove(a)
      } else {
         return AiMove(a, d, e)
      }
}

func HumanMove(a board.Board) int {
      for i := 0; i < len(a.Array); i++ {
          if a.Array[i] == strconv.Itoa(i+1) {
	     return i
          }
      }
      return 0
}

func AiMove(a board.Board, d player.Player, e player.Player) int {
    availableSpots := a.GetAvailableSpots()
    var score, bestMove int
    bestScore := -100
    depth := 1
    
    boardCopy := make([]string, len(a.Array))
    copy(boardCopy, a.Array)      
    currentBoardState := a

    for i := 0; i < len(availableSpots); i++ {
      currentBoardState.MakeMove(availableSpots[i], d.Symbol)
      score = Minimax(currentBoardState, d, e, depth + 1)
      if score > bestScore {
        bestScore = score
        bestMove = availableSpots[i]
      }
      currentBoardState.Array = make([]string, len(a.Array))
      copy(currentBoardState.Array, boardCopy)
   }
    return bestMove	 
}

func Minimax(a board.Board, d player.Player, e player.Player, depth int) int {
     if a.WinningCombination() || a.TieCombination() {
       winner := d.Symbol
       return Score(a, d, winner, depth)
     }
  
     var currentPlayer player.Player
     var score int
     leastScore := 100
     availableSpots := a.GetAvailableSpots()
     
     boardCopy := make([]string, len(a.Array))
     copy(boardCopy, a.Array)      
     currentBoardState := a
      
     for i := 0; i < len(availableSpots); i++ {
        currentPlayer = d.SwitchPlayer(d, e)
        currentBoardState.MakeMove(availableSpots[i], currentPlayer.Symbol)
        score = -Minimax(currentBoardState, currentPlayer, d, depth + 1)
        if score < leastScore {
           leastScore = score
        }
        currentBoardState.Array = make([]string, len(a.Array))
        copy(currentBoardState.Array, boardCopy)
      }
      return leastScore
}

func Score(a board.Board, d player.Player, winner string, depth int) int {
     if a.WinningCombination() && winner == d.Symbol {
        return 10/depth
     } else if a.WinningCombination() && winner != d.Symbol {
        return -10/depth
     } else if a.TieCombination() {
        return 0
     } else {
        return 0
     }
}

