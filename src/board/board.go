package board

import "errors"
import "strconv"

type Board struct {
  Size, AvailableSpots int
  Array []string
}

func (board Board) Create() Board {
    size := board.Size
    board.AvailableSpots = (size*size)
    board.Array = board.FillArray(size*size)
    return board
}

func (board Board) FillArray(length int) []string {
     board.Array = make([]string, length)
     for i := 0; i < length; i++ {
        board.Array[i] = strconv.Itoa(i + 1)
     }
     return board.Array
}

func (board Board) MakeMove(position int, symbol string) (Board, error) {
    boardLength := (board.Size * board.Size)

    if position < 0 || position > boardLength - 1 {
        return board, errors.New("position is Out of Range")
    } else if board.Array[position] != strconv.Itoa(position + 1) {
        return board, errors.New("position already filled")
    }
    
    board.Array[position] = symbol
    board.AvailableSpots--
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

