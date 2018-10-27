package play

import (
        ."../../../core/src/board"
        ."../../../core/src/game"
        ."../../../core/src/moveStrategy"
        ."../../../core/src/player"
        ."../setup"
        ."fmt"
)

func StartGame() {
     PrintStartOptions()

     var choice int
     Scan(&choice)
     
     PlayOrQuit(choice)
}

func PrintStartOptions() {
     Println("\n")
     Println("1. Start Game")
     Println("2. Quit")
}

func PlayOrQuit(choice int) {     
     switch choice {
     case 1:
        Println("--------")   
        board, firstPlayer, otherPlayer := SetUp()
        PrintBoard(board)
        Play(board, firstPlayer, otherPlayer)
     case 2:
        Println("GoodBye...")
     default:
        StartGame()
     }
}

func PrintBoard(b Board) {
   Println("------")
   Println(b.Array[0] + " " + b.Array[1] + " " + b.Array[2] )
   Println(b.Array[3] + " " + b.Array[4] + " " + b.Array[5] )
   Println(b.Array[6] + " " + b.Array[7] + " " + b.Array[8] )
   Println("------")
}

func Play(b Board, firstPlayer Player, otherPlayer Player) {
    position := GetMove(b, firstPlayer, otherPlayer)
    c := validateMove(b, position, firstPlayer)
    PrintBoard(c)
    CheckGameStatus(c, firstPlayer, otherPlayer)
}

func GetMove(b Board, firstPlayer Player, otherPlayer Player) int { 
    var position int
    switch firstPlayer.Type {
    case HUMAN:
       position = GetHumanMove()
    default:
       Println("Ai Move")
       position = AiMove(b, firstPlayer, otherPlayer)
    }
    return position
}

func GetHumanMove() int {
    Print("\nEnter Move :")
    var move int
    Scan(&move)
    actualMove := move - 1
    return actualMove
}

func validateMove(b Board, position int, firstPlayer Player) Board {
    c, err := b.MakeMove(position, firstPlayer.Symbol)
    for err != nil {
       Print(err)
       position = GetHumanMove()
       c, err = b.MakeMove(position, firstPlayer.Symbol)
    }
    return c
}

func CheckGameStatus(b Board, firstPlayer Player, otherPlayer Player) {
    status := Status(b, firstPlayer)
    switch status {
    case NO_WINNER_YET:
       currentPlayer := firstPlayer.SwitchPlayer(firstPlayer, otherPlayer)
       otherPlayer = firstPlayer
       Play(b, currentPlayer, otherPlayer)
    default:
       Print(status)
       StartGame()
    }
}
