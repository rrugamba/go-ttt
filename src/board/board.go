package board

import "errors"

type Board struct {
  Size, AvailableSpots int
  Array []string
}

func (board Board) Create() []string {
    size := board.Size
    board.Array = make([]string, (size*size))
    return board.Array
}

func (board Board) MakeMove(position int, symbol string) (Board, error) {
    boardLength := (board.Size * board.Size)
    if position < 1 || position > boardLength {
        return board, errors.New("position is Out of Range")
    }
    boardArray := board.Create()
    boardArray[position] = symbol
    board.AvailableSpots--
    board.Array = boardArray
    return board, nil
}
