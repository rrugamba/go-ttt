package moveStrategy

import (
        ."../board"
        ."../player"
)

func AiMove(a Board, d Player, e Player) int {
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

    copy(a.Array, boardCopy)  
    return bestMove	 
}

func Minimax(a Board, d Player, e Player, depth int) int {
     if a.WinningCombination() || a.TieCombination() {
       winner := d.Symbol
       return Score(a, d, winner, depth)
     }
  
     var currentPlayer Player
     var score int
     leastScore := 100
     availableSpots := a.GetAvailableSpots()
     
     boardCopy := make([]string, len(a.Array))
     copy(boardCopy, a.Array)      
     currentBoardState := a
     currentPlayer = d.SwitchPlayer(d, e)
     
     for i := 0; i < len(availableSpots); i++ {
        currentBoardState.MakeMove(availableSpots[i], currentPlayer.Symbol)
        score = -Minimax(currentBoardState, currentPlayer, d, depth + 1)
        if score < leastScore {
           leastScore = score
        }
        currentBoardState.Array = make([]string, len(a.Array))
        copy(currentBoardState.Array, boardCopy)
      }
 
    copy(a.Array, boardCopy)  
    return leastScore
}

func Score(a Board, d Player, winner string, depth int) int {
     if a.WinningCombination() && winner == d.Symbol {
        return 10/depth
     } else if a.WinningCombination() && winner != d.Symbol {
        return -10/depth
     } else {
        return 0
     } 
}

