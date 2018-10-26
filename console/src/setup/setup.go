package setup

import (
        "../../../core/src/board"
        "../../../core/src/player"
        "fmt"
        "bufio"
        "os"
        "strings"
)

func GetPlayerSymbol() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\nSymbol: ")
    symbol, _ := reader.ReadString('\n')
    symbol = strings.TrimSuffix(symbol, "\n")
    return strings.ToUpper(symbol)    
}

func GetPlayerType() string {
    fmt.Println("1. human")
    fmt.Println("2. ai")
    var choosenValue int
    fmt.Scan(&choosenValue)
    return Type(choosenValue)
}

func Type(input int) string {
     switch input {
     case 1:
         return "human"
     case 2:
         return "ai"
     default:
         return GetPlayerType()
     }
}
 
func GetOtherPlayer(p player.Player) player.Player {
    fmt.Print("Second Player \n")
    symbol := GetPlayerSymbol()
   
    for symbol == p.Symbol {
       symbol = GetPlayerSymbol()
    }
    
    switch p.Type {
    case "human": 
       return player.Player {Symbol: symbol, Type: Type(2),}
    default:
       return player.Player {Symbol: symbol, Type: Type(1),}
    }
}

func WhoPlaysFirst(p1 player.Player, p2 player.Player) player.Player {
    fmt.Println("Who Plays First? human or ai?")
    fmt.Println("1. human")
    fmt.Println("2. ai")
    var choosenValue int
    fmt.Scan(&choosenValue)
    return GetFirstPlayer(p1, p2, choosenValue)
}

func GetFirstPlayer(p1 player.Player, p2 player.Player, input int) player.Player {
    if (input == 1 && p1.Type == "human") || (input == 2 && p1.Type == "ai") {
       return p1
    } else if (input == 1 && p2.Type == "human") || (input == 2 && p2.Type == "ai") { 
       return p2
    }  else {
       return WhoPlaysFirst(p1, p2)
    }
}

func SetUp() (board.Board, player.Player, player.Player) {
     b := board.Board {Size: 3,}
     newBoard := b.Create()
 
    fmt.Println("Player 1: ")
    s := GetPlayerSymbol()

    typo := GetPlayerType()

    player1 := player.Player{Symbol: s, Type: typo,}
    player2 := GetOtherPlayer(player1)

    firstPlayer := WhoPlaysFirst(player1, player2)
    otherPlayer := firstPlayer.SwitchPlayer(player1, player2) 

    return newBoard, firstPlayer, otherPlayer
}
