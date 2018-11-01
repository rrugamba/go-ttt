package play

import (
        ."../../../core/src/board"
        ."../../../core/src/game"
        ."../../../core/src/moveStrategy"
        ."../../../core/src/player"
        ."../setup"
        ."fmt"
        "time"
)

func StartGame() {
     printStartOptions()

     var choice int
     Scan(&choice)
     
     playOrQuit(choice)
}

func printStartOptions() {
     Println("\n")
     Println("1. Start Game")
     Println("2. Quit")
}

func playOrQuit(choice int) {     
     switch choice {
     case 1:
        Println("--------")   
        board, firstPlayer, otherPlayer := SetUp()
        PrintBoard(board)
        play(board, firstPlayer, otherPlayer)
     case 2:
        Println("GoodBye...")
     default:
        StartGame()
     }
}

func PrintBoard(b Board) {
   Println("-------")
   Println(b.Array[0] + "  " + b.Array[1] + "  " + b.Array[2]) 
   Println(b.Array[3] + "  " + b.Array[4] + "  " + b.Array[5])
   Println(b.Array[6] + "  " + b.Array[7] + "  " + b.Array[8]) 
   Println("-------")
}

func play(b Board, firstPlayer Player, otherPlayer Player) {
    position := getMove(b, firstPlayer, otherPlayer)
    c := validateMove(b, position, firstPlayer)
    PrintBoard(c)
    checkGameStatus(c, firstPlayer, otherPlayer)
}

func getMove(b Board, firstPlayer Player, otherPlayer Player) int { 
    var position int
    switch firstPlayer.Type {
    case HUMAN:
       position = getHumanMove()
    default:
       Println("Ai Move")
       time.Sleep(2000 * time.Millisecond)
       position = AiMove(b, firstPlayer, otherPlayer)
    }
    return position
}

func getHumanMove() int {
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
       position = getHumanMove()
       c, err = b.MakeMove(position, firstPlayer.Symbol)
    }
    return c
}

func checkGameStatus(b Board, firstPlayer Player, otherPlayer Player) {
    status := Status(b, firstPlayer)
    switch status {
    case NO_WINNER_YET:
       currentPlayer := firstPlayer.SwitchPlayer(firstPlayer, otherPlayer)
       otherPlayer = firstPlayer
       play(b, currentPlayer, otherPlayer)
    default:
       Print(status)
       StartGame()
    }
}
