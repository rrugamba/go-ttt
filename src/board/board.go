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

    if position < 1 || position > boardLength {
        return board, errors.New("position is Out of Range")
    } else if board.Array[position] != "" {
        return board, errors.New("position already filled")
    }
    
    board.Array[position] = symbol
    board.AvailableSpots--
    return board, nil
}

func (board Board) HorizontalWinningCombination() bool {
    return false
}


