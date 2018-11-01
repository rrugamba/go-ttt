package setup

import (
        ."../../../core/src/board"
        ."../../../core/src/player"
        ."fmt"
        "bufio"
        ."os"
        "strings"
        ."strconv"
)

func getPlayerSymbol() string {
    reader := bufio.NewReader(Stdin)
    Print("Symbol: ")
    symbol, _ := reader.ReadString('\n')
    symbol = strings.TrimSuffix(symbol, "\n")
    s := checkIfNumber(symbol)
    return strings.ToUpper(s)    
}

func checkIfNumber(symbol string) string {
   if _, err := Atoi(symbol); err == nil {
      Println("enter anything thats not a number.")
      return getPlayerSymbol()
   }
   return symbol
}

func getPlayerType() string {
    Print("Type: ")
    Print("(1) human. ")
    Print("(2) ai.  :enter 1 or 2:  ")
    var choosenValue int
    Scan(&choosenValue)
    Println("------")
    return whichType(choosenValue)
}

func whichType(input int) string {
     switch input {
     case 1:
         return HUMAN
     case 2:
         return AI
     default:
         return getPlayerType()
     }
}
 
func getOtherPlayer(player Player) Player {
    Print("Player 2 \n")
    symbol := getPlayerSymbol()
   
    for symbol == player.Symbol {
       Println("Use different symbol from Player 1, please!")
       symbol = getPlayerSymbol()
    }
    
    switch player.Type {
    case HUMAN: 
       return Player {Symbol: symbol, Type: whichType(2),}
    default:
       return Player {Symbol: symbol, Type: whichType(1),}
    }
}

func whoPlaysFirst(p1 Player, p2 Player) Player {
    Print("Who Plays First? (1) human. (2) ai.  :enter 1 or 2: ")
    var choosenValue int
    Scan(&choosenValue)
    return getFirstPlayer(p1, p2, choosenValue)
}

func getFirstPlayer(p1 Player, p2 Player, input int) Player {
    if (input == 1 && p1.Type == HUMAN) || (input == 2 && p1.Type == AI) {
       return p1
    } else if (input == 1 && p2.Type == HUMAN) || (input == 2 && p2.Type == AI) { 
       return p2
    }  else {
       return whoPlaysFirst(p1, p2)
    }
}

func SetUp() (Board, Player, Player) {
    b := Board {Size: 3,}
    newBoard := b.Create()
 
    Println("Player 1 ")
    symbol := getPlayerSymbol()

    pType := getPlayerType()

    player1 := Player{Symbol: symbol, Type: pType,}
    player2 := getOtherPlayer(player1)
    Println("Type: " + player2.Type)
    Println("------")
 
    firstPlayer := whoPlaysFirst(player1, player2)
    otherPlayer := firstPlayer.SwitchPlayer(player1, player2) 

    return newBoard, firstPlayer, otherPlayer
}
