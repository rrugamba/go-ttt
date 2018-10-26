package player_test

import (
         "testing"
         ."github.com/franela/goblin"
         "../../src/player"
)

func Test(t *testing.T) {
     g := Goblin(t)
     g.Describe("Creating a Player", func() {
         g.It("should get the symbol and type of player", func() {
             newPlayer := player.Player {
               Symbol: "X",
               Type: "human",
             }
             g.Assert(newPlayer.GetSymbol()).Equal("X")
             g.Assert(newPlayer.GetType()).Equal("human")
         })

        g.It("should get the symbol and type of other player", func() {
             playerOne := player.Player {
               Symbol: "X",
               Type: "human",
             }

             playerTwo := player.Player {
               Symbol: "O",
               Type: "ai",
             }
             
             g.Assert(playerOne.SwitchPlayer(playerOne, playerTwo)).Equal(playerTwo)
             g.Assert(playerOne.SwitchPlayer(playerOne, playerTwo).Symbol).Equal("O")
             g.Assert(playerOne.SwitchPlayer(playerOne, playerTwo).Type).Equal("ai")
         })
        g.It("should get the symbol and type of other player", func() {
             playerOne := player.Player {
               Symbol: "X",
               Type: "human",
             }

             playerTwo := player.Player {
               Symbol: "O",
               Type: "ai",
             }
             
             g.Assert(playerTwo.SwitchPlayer(playerOne, playerTwo)).Equal(playerOne)
             g.Assert(playerTwo.SwitchPlayer(playerOne, playerTwo).Symbol).Equal("X")
             g.Assert(playerTwo.SwitchPlayer(playerOne, playerTwo).Type).Equal("human")
         })
     })
}


