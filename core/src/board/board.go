package board

import (
        ."errors"
        ."strconv"
)

const OUT_OF_RANGE_ERROR = "position is Out of Range"
const POSITION_ALREADY_FILLED_ERROR = "position already filled"

type Board struct {
  Size int
  Array []string
}

func (board Board) Create() Board {
    size := board.Size
    board.Array = board.FillArray(size*size)
    return board
}

func (board Board) FillArray(length int) []string {
     board.Array = make([]string, length)
     for i := 0; i < length; i++ {
        board.Array[i] = Itoa(i + 1)
     }
     return board.Array
}

func (board Board) GetAvailableSpots() []int {
     var arrayOfAvailableSpots []int
     arrayLength := len(board.Array)
     for i := 0; i < arrayLength; i++ {
        if board.Array[i] == Itoa(i + 1) { 
           arrayOfAvailableSpots = append(arrayOfAvailableSpots, i) 
        }
      }
     return arrayOfAvailableSpots
}

func (board Board) MakeMove(position int, symbol string) (Board, error) {
    boardLength := (board.Size * board.Size)

    if position < 0 || position > boardLength - 1 {
        return board, New(OUT_OF_RANGE_ERROR)
    } else if board.Array[position] != Itoa(position + 1) {
        return board, New(POSITION_ALREADY_FILLED_ERROR)
    }
    
    board.Array[position] = symbol
    return board, nil
}

func (board Board) WinningCombination() bool {
     return board.HorizontalWinningCombination() ||
            board.VerticalWinningCombination() ||
            board.DiagonalWinningCombination()
}

func (board Board) HorizontalWinningCombination() bool {
      return (board.Array[0] == board.Array[1] && board.Array[1] == board.Array[2]) ||
             (board.Array[3] == board.Array[4] && board.Array[4] == board.Array[5]) ||
             (board.Array[6] == board.Array[7] && board.Array[7] == board.Array[8])
}

func (board Board) VerticalWinningCombination() bool {
      return (board.Array[0] == board.Array[3] && board.Array[3] == board.Array[6]) ||
             (board.Array[1] == board.Array[4] && board.Array[4] == board.Array[7]) ||
             (board.Array[2] == board.Array[5] && board.Array[5] == board.Array[8])
}

func (board  Board) DiagonalWinningCombination() bool {
      return (board.Array[0] == board.Array[4] && board.Array[4] == board.Array[8]) ||
             (board.Array[2] == board.Array[4] && board.Array[4] == board.Array[6]) 
}

func (board Board) TieCombination() bool {
     return !board.WinningCombination() && len(board.GetAvailableSpots()) == 0
}
