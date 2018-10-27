package mockConsole

import  ."../../../core/src/player"

func MockGetPlayerType() string {
   return "human"
}

func MockGetPlayerSymbol() string {
   return "O"
}

func MockGetOtherPlayer(p Player) Player {
   symbol := "X"
   if p.Symbol == symbol {
      symbol = "Y"
   }
   var typo string
   if p.Type == "human" {
      typo = "ai"
   } else {
      typo = "human"
   }
  return Player {Symbol: symbol, Type: typo,}
}
