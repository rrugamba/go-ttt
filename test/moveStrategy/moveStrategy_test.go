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
  
          p1 := player.Player{Symbol: "X", Type: "human",}
          p2 := player.Player{Symbol: "O", Type: "ai",}
          
          g.Assert(moveStrategy.GetMove(b, p1, p2)).Equal(0)
        })
     
        g.It("should make move(2) as best move on a pre-filled board", func() {
        
          newBoard := board.Board{Size: 3,}
          b := newBoard.Create()
          c, _ := b.MakeMove(0, "X")
          d, _ := c.MakeMove(1, "X")
  
          p1 := player.Player{Symbol: "X", Type: "human",}
          p2 := player.Player{Symbol: "O", Type: "ai",}
          
          g.Assert(moveStrategy.GetMove(d, p1, p2)).Equal(2)
        })
    })
    g.Describe ("Move strategy for Ai Player", func(){
        g.It("ai WIN MOVE: should make move(2) as best move on current board state", func() {

         newBoard := board.Board{Size: 3,}
         b := newBoard.Create()
 
         p1 := player.Player{Symbol: "O", Type: "ai",}
         p2 := player.Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         d, _ := c.MakeMove(1, "O")
         e, _ := d.MakeMove(3, "X")

         g.Assert(moveStrategy.GetMove(e, p1, p2)).Equal(2)
       })
     
      g.It("ai WIN MOVE: should make move(6) as best move on current board state", func() {

         newBoard := board.Board{Size: 3,}
         b := newBoard.Create()
 
         p1 := player.Player{Symbol: "O", Type: "ai",}
         p2 := player.Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         d, _ := c.MakeMove(1, "X")
         e, _ := d.MakeMove(4, "X")
         f, _ := e.MakeMove(3, "O")
         g.Assert(moveStrategy.GetMove(f, p1, p2)).Equal(6)
       })

       g.It("ai BLOCK MOVE: should make move(7) as best move on current board state", func() {

         newBoard := board.Board{Size: 3,}
         b := newBoard.Create()
 
         p1 := player.Player{Symbol: "O", Type: "ai",}
         p2 := player.Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(2, "O")
         d, _ := i.MakeMove(1, "X")
         e, _ := d.MakeMove(4, "X")
         f, _ := e.MakeMove(3, "O")
         k, _ := f.MakeMove(6, "X")
         g.Assert(moveStrategy.GetMove(k, p1, p2)).Equal(7)
       })
      g.It("ai WIN MOVE: should make move(4) as best move on current board state", func() {

         newBoard := board.Board{Size: 3,}
         b := newBoard.Create()
 
         p1 := player.Player{Symbol: "O", Type: "ai",}
         p2 := player.Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "X")
         i, _ := c.MakeMove(1, "X")
         d, _ := i.MakeMove(2, "O")
         e, _ := d.MakeMove(3, "O")
         f, _ := e.MakeMove(5, "X")
         k, _ := f.MakeMove(7, "O")
         g.Assert(moveStrategy.GetMove(k, p1, p2)).Equal(4)
       })
     g.It("ai BLOCK MOVE: should make move(7) as best move on current board state", func() {

         newBoard := board.Board{Size: 3,}
         b := newBoard.Create()
 
         p1 := player.Player{Symbol: "O", Type: "ai",}
         p2 := player.Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(1, "X")
         d, _ := i.MakeMove(2, "O")
         e, _ := d.MakeMove(4, "X")
         g.Assert(moveStrategy.GetMove(e, p1, p2)).Equal(7)
       })
   })

}


