package play

import (
        "../../../core/src/board"
        "../../../core/src/game"
        "../../../core/src/moveStrategy"
        "../../../core/src/player"
        "../setup"
        "fmt"
)

func GetHumanMove() int {
    fmt.Print("\nEnter Move ")
    var move int
    fmt.Scan(&move)
    actualMove := move - 1
    return actualMove
}

func PrintBoard(b board.Board) {
   fmt.Println(b.Array[0] + " " + b.Array[1] + " " + b.Array[2] )
   fmt.Println(b.Array[3] + " " + b.Array[4] + " " + b.Array[5] )
   fmt.Println(b.Array[6] + " " + b.Array[7] + " " + b.Array[8] )
}

func Play(b board.Board, firstPlayer player.Player, otherPlayer player.Player) {
    var position int
    if firstPlayer.Type == "human" {
       position = GetHumanMove()
    } else {
       fmt.Println("Ai Move \n")
       boardCopy := make([]string, len(b.Array))
       copy(boardCopy, b.Array)

       position = moveStrategy.AiMove(b, firstPlayer, otherPlayer)

       b.Array = make([]string, len(b.Array))
       copy(b.Array, boardCopy)
    }

    c, err := b.MakeMove(position, firstPlayer.Symbol)
    for err != nil {
       fmt.Print(err)
       position = GetHumanMove()
       c, err = b.MakeMove(position, firstPlayer.Symbol)
    }
    status := game.Status(c, firstPlayer)
    PrintBoard(c)
    fmt.Println("-------")
    
    switch status {
    case "NO WINNER YET":
       currentPlayer := firstPlayer.SwitchPlayer(firstPlayer, otherPlayer)
       otherPlayer = firstPlayer
       Play(c, currentPlayer, otherPlayer)
    default:
       fmt.Print(status)
       StartGame()
    }
}

func PrintStartOptions() {
     fmt.Println("\n")
     fmt.Println("1. Start Game")
     fmt.Println("2. Quit")
}

func StartGame() {
     PrintStartOptions()

     var choice int
     fmt.Scan(&choice)
     
     switch choice {
     case 1:
        fmt.Println("--------")   
        board, firstPlayer, otherPlayer := setup.SetUp()
        PrintBoard(board)
        Play(board, firstPlayer, otherPlayer)
     case 2:
        fmt.Println("GoodBye...")
     default:
        StartGame()
     }
}
