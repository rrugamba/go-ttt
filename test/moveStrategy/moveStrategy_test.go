package test_moveStrategy

import (
        "testing"
        ."github.com/franela/goblin"
        "../../src/moveStrategy"
        "../../src/board"
        "../../src/player"
)


func Test(t *testing.T) {
     g := Goblin(t)
     g.Describe ("Move strategy for Human a Player", func(){
         g.It("should make move(0) as best move on an empty board", func() {

          newBoard := board.Board{Size: 3,}
          b := newBoard.Create()
  
          p := player.Player{Symbol: "X", Type: "human",}

          g.Assert(moveStrategy.GetMove(b, p)).Equal(0)
        })
     
        g.It("should make move(2) as best move on a pre-filled board", func() {
        
          newBoard := board.Board{Size: 3,}
          b := newBoard.Create()
          c, _ := b.MakeMove(0, "X")
          d, _ := c.MakeMove(1, "X")
  
          p := player.Player{Symbol: "X", Type: "human",}
          
          g.Assert(moveStrategy.GetMove(d, p)).Equal(2)
        })
    })
    g.Describe ("Move strategy for Ai Player", func(){
        g.It("should make move(2) as best move on current board state", func() {

         newBoard := board.Board{Size: 3,}
         b := newBoard.Create()
 
         p := player.Player{Symbol: "O", Type: "ai",}

         c, _ := b.MakeMove(0, "O")
         d, _ := c.MakeMove(1, "O")
         e, _ := d.MakeMove(3, "X")

         g.Assert(moveStrategy.GetMove(e, p)).Equal(2)
       })
     
   })

}


