package test_setup

import (
        "testing"
        . "github.com/franela/goblin"
        ."../../../core/src/board"
        ."../../../core/src/player"
        ."../mockConsole"
)

func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe ("Creating a Board of any size", func() {
        g.It("should create board size of 3", func(){
          b := Board {Size: 3,}
          newBoard := b.Create()
          g.Assert(len(newBoard.Array)).Equal(9)
        })

        g.It("should get symbol and type of player", func(){
          typ := MockGetPlayerType() 
          symbol := MockGetPlayerSymbol()
          player1 := Player{Symbol: symbol, Type: typ,}
         
         g.Assert(player1.Symbol).Equal("O")
         g.Assert(player1.Type).Equal("human")
        })

        g.It("should get symbol and type of second player", func(){
            typ := MockGetPlayerType() 
            symbol := MockGetPlayerSymbol()
            player1 := Player{Symbol: symbol, Type: typ,}
         g.Assert(player1.Symbol).Equal("O")
         g.Assert(player1.Type).Equal("human")

            player2 := Player{Symbol: "X", Type: "ai"}
         g.Assert(MockGetOtherPlayer(player1)).Equal(player2)
        })

    })
}

