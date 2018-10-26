package test_setup

import (
        "testing"
        . "github.com/franela/goblin"
        "../../../core/src/board"
        "../../../core/src/player"
)

func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe ("Creating a Board of any size", func() {
        g.It("should create board size of 3", func(){
          b := board.Board {Size: 3,}
          newBoard := b.Create()
          g.Assert(len(newBoard.Array)).Equal(9)
        })

        g.It("should get symbol and type of player", func(){
          typ := MockGetPlayerType() 
          symbol := MockGetPlayerSymbol()
          player1 := player.Player{Symbol: symbol, Type: typ,}
         
         g.Assert(player1.Symbol).Equal("O")
         g.Assert(player1.Type).Equal("human")
        })

        g.It("should get symbol and type of second player", func(){
            typ := MockGetPlayerType() 
            symbol := MockGetPlayerSymbol()
            player1 := player.Player{Symbol: symbol, Type: typ,}
         g.Assert(player1.Symbol).Equal("O")
         g.Assert(player1.Type).Equal("human")

            player2 := player.Player{Symbol: "X", Type: "ai"}
         g.Assert(MockGetOtherPlayer(player1)).Equal(player2)
        })

    })
}

func MockGetPlayerType() string {
   return "human"
}

func MockGetPlayerSymbol() string {
   return "O"
}

func MockGetOtherPlayer(p player.Player) player.Player {
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
  return player.Player {Symbol: symbol, Type: typo,}
}
