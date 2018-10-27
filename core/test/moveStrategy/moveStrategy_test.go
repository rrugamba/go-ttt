package moveStrategy_test

import (
        "testing"
        ."github.com/franela/goblin"
        ."../../src/moveStrategy"
        ."../../src/board"
        ."../../src/player"
)


func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe ("Move strategy for Ai Player", func(){
     g.It("ai TIE MOVE: should make move(8) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(1, "X")
         t, _ := i.MakeMove(2, "O")
         d, _ := t.MakeMove(3, "O")
         k, _ := d.MakeMove(4, "X")
         e, _ := k.MakeMove(5, "X")
         o, _ := e.MakeMove(6, "X")
         q, _ := o.MakeMove(7, "O")
         g.Assert(AiMove(q, p1, p2)).Equal(8)
       })
     g.It("ai BLOCK MOVE: should make move(7) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(1, "X")
         t, _ := i.MakeMove(2, "O")
         d, _ := t.MakeMove(3, "O")
         k, _ := d.MakeMove(4, "X")
         e, _ := k.MakeMove(5, "X")
         o, _ := e.MakeMove(6, "X")
         g.Assert(AiMove(o, p1, p2)).Equal(7)
       })
        g.It("ai WIN MOVE: should make move(2) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         d, _ := c.MakeMove(1, "O")
         e, _ := d.MakeMove(3, "X")

         g.Assert(AiMove(e, p1, p2)).Equal(2)
       })
       g.It("ai BLOCK MOVE: should make move(8) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "X")
         d, _ := c.MakeMove(3, "O")
         e, _ := d.MakeMove(4, "X")

         g.Assert(AiMove(e, p1, p2)).Equal(8)
       })
      g.It("ai WIN MOVE: should make move(6) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         d, _ := c.MakeMove(1, "X")
         e, _ := d.MakeMove(4, "X")
         f, _ := e.MakeMove(3, "O")
         g.Assert(AiMove(f, p1, p2)).Equal(6)
       })

      g.It("ai BLOCK MOVE: should make move(7) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         d, _ := c.MakeMove(1, "X")
         e, _ := d.MakeMove(4, "X")
         f, _ := e.MakeMove(5, "O")
         g.Assert(AiMove(f, p1, p2)).Equal(7)
       })
       g.It("ai BLOCK MOVE: should make move(7) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(2, "O")
         d, _ := i.MakeMove(1, "X")
         e, _ := d.MakeMove(4, "X")
         f, _ := e.MakeMove(3, "O")
         k, _ := f.MakeMove(6, "X")
         g.Assert(AiMove(k, p1, p2)).Equal(7)
       })
      g.It("ai WIN MOVE: should make move(4) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "X")
         i, _ := c.MakeMove(1, "X")
         d, _ := i.MakeMove(2, "O")
         e, _ := d.MakeMove(3, "O")
         f, _ := e.MakeMove(5, "X")
         k, _ := f.MakeMove(6, "O")
         g.Assert(AiMove(k, p1, p2)).Equal(4)
       })
     g.It("ai BLOCK MOVE: should make move(7) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(1, "X")
         d, _ := i.MakeMove(2, "O")
         e, _ := d.MakeMove(4, "X")
         g.Assert(AiMove(e, p1, p2)).Equal(7)
       })
     g.It("ai WIN MOVE: should make move(5) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "X")
         i, _ := c.MakeMove(1, "O")
         d, _ := i.MakeMove(3, "O")
         k, _ := d.MakeMove(7, "X")
         e, _ := k.MakeMove(4, "O")
         g.Assert(AiMove(e, p1, p2)).Equal(5)
       })
     g.It("ai WIN MOVE: should make move(6) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "X")
         i, _ := c.MakeMove(1, "O")
         z, _ := i.MakeMove(2, "O")
         d, _ := z.MakeMove(3, "X")
         k, _ := d.MakeMove(4, "O")
         e, _ := k.MakeMove(8, "X")
         g.Assert(AiMove(e, p1, p2)).Equal(6)
       })
     g.It("ai BLOCK MOVE: should make move(5) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(1, "X")
         t, _ := i.MakeMove(2, "O")
         d, _ := t.MakeMove(3, "X")
         k, _ := d.MakeMove(7, "O")
         e, _ := k.MakeMove(4, "X")
         g.Assert(AiMove(e, p1, p2)).Equal(5)
       })
     
     g.It("ai BLOCK MOVE: should make move(5) as best move on current board state", func() {

         board := Board{Size: 3,}
         b := board.Create()
 
         p1 := Player{Symbol: "O", Type: "ai",}
         p2 := Player{Symbol: "X", Type: "human",}

         c, _ := b.MakeMove(0, "O")
         i, _ := c.MakeMove(1, "X")
         d, _ := i.MakeMove(3, "X")
         k, _ := d.MakeMove(7, "O")
         e, _ := k.MakeMove(4, "X")
         g.Assert(AiMove(e, p1, p2)).Equal(5)
       })
   })

}


