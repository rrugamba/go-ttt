package test_other

import (
        "testing"
        ."github.com/franela/goblin"
        "../../src/other"
        "../../src/player"   
     )


func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe(" Switching players, getting 'other' player", func() {
       g.It("get other playerB, given a current playerA", func() {
          playerA := player.Player{Symbol: "X", Type: "human",}
          playerB := player.Player{Symbol: "O", Type: "ai",}
         
          array := make([] player.Player, 2)
          array[0] = playerA
          array[1] = playerB

          players := other.Players {
	      Array: array,
          }
          currentPlayer := playerA
          otherPlayer := players.SwitchPlayers(currentPlayer)
	  g.Assert(otherPlayer.Type).Equal("ai")
      })

       g.It("get other playerA, given a current playerB", func() {
          playerA := player.Player{Symbol: "X", Type: "human",}
          playerB := player.Player{Symbol: "O", Type: "ai",}
         
          array := make([] player.Player, 2)
          array[0] = playerA
          array[1] = playerB

          players := other.Players {
	      Array: array,
          }
          currentPlayer := playerB
          otherPlayer := players.SwitchPlayers(currentPlayer)
	  g.Assert(otherPlayer.Type).Equal("human")
      })
   }) 
}
