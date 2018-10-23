package test_player

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
     })
}


