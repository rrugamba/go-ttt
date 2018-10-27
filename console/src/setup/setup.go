package setup

import (
        ."../../../core/src/board"
        ."../../../core/src/player"
        ."fmt"
        "bufio"
        ."os"
        "strings"
)

func GetPlayerSymbol() string {
    reader := bufio.NewReader(Stdin)
    Print("Symbol: ")
    symbol, _ := reader.ReadString('\n')
    symbol = strings.TrimSuffix(symbol, "\n")
    return strings.ToUpper(symbol)    
}

func GetPlayerType() string {
    Print("Type: ")
    Print("(1) human. ")
    Print("(2) ai. :")
    var choosenValue int
    Scan(&choosenValue)
    Println("------")
    return Type(choosenValue)
}

func Type(input int) string {
     switch input {
     case 1:
         return HUMAN
     case 2:
         return AI
     default:
         return GetPlayerType()
     }
}
 
func GetOtherPlayer(player Player) Player {
    Print("Player 2 \n")
    symbol := GetPlayerSymbol()
   
    for symbol == player.Symbol {
       symbol = GetPlayerSymbol()
    }
    
    switch player.Type {
    case HUMAN: 
       return Player {Symbol: symbol, Type: Type(2),}
    default:
       return Player {Symbol: symbol, Type: Type(1),}
    }
}

func WhoPlaysFirst(p1 Player, p2 Player) Player {
    Println("Who Plays First? 1 or 2?")
    Println("1. human")
    Print("2. ai : ")
    var choosenValue int
    Scan(&choosenValue)
    return GetFirstPlayer(p1, p2, choosenValue)
}

func GetFirstPlayer(p1 Player, p2 Player, input int) Player {
    if (input == 1 && p1.Type == HUMAN) || (input == 2 && p1.Type == AI) {
       return p1
    } else if (input == 1 && p2.Type == HUMAN) || (input == 2 && p2.Type == AI) { 
       return p2
    }  else {
       return WhoPlaysFirst(p1, p2)
    }
}

func SetUp() (Board, Player, Player) {
    b := Board {Size: 3,}
    newBoard := b.Create()
 
    Println("Player 1: ")
    symbol := GetPlayerSymbol()

    pType := GetPlayerType()

    player1 := Player{Symbol: symbol, Type: pType,}
    player2 := GetOtherPlayer(player1)
    Println("Type: " + player2.Type)
    Println("------")
 
    firstPlayer := WhoPlaysFirst(player1, player2)
    otherPlayer := firstPlayer.SwitchPlayer(player1, player2) 

    return newBoard, firstPlayer, otherPlayer
}
