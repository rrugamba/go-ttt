package board

import "errors"

type Board struct {
  Size, AvailableSpots int
  Array []string
}

func (board Board) Create() Board {
    size := board.Size
    board.AvailableSpots = (size*size)
    board.Array = make([]string, (size*size))
    return board
}

func (board Board) MakeMove(position int, symbol string) (Board, error) {
    boardLength := (board.Size * board.Size)

    if position < 0 || position > boardLength - 1 {
        return board, errors.New("position is Out of Range")
    } else if board.Array[position] != "" {
        return board, errors.New("position already filled")
    }
    
    board.Array[position] = symbol
    board.AvailableSpots--
    return board, nil
}

func (board Board) HorizontalWinningCombination() bool {
   boardLength := (board.Size * board.Size)
   if board.AvailableSpots < boardLength { 
      if board.Array[0] == board.Array[1] && board.Array[1] == board.Array[2] {
          return true
       }
   }
   return false
}

func (board  Board) VerticalWinningCombination() bool {
   boardLength := (board.Size * board.Size)
   if board.AvailableSpots < boardLength { 
      if board.Array[0] == board.Array[3] && board.Array[3] == board.Array[6] {
          return true
       }
   }
    return false
}

func (board  Board) DiagonalWinningCombination() bool {
   boardLength := (board.Size * board.Size)
   if board.AvailableSpots < boardLength { 
      if board.Array[0] == board.Array[4] && board.Array[4] == board.Array[8] {
          return true
       }
   }
    return false
}

