package game_test

import (
         "testing"
         ."github.com/franela/goblin"
         ."../../src/game"
         ."../../src/board"
         ."../../src/player"
)

func Test(t *testing.T) {
      g := Goblin(t)
      g.Describe("Game set up", func() {
         g.It("one move on board, no winner yet", func() {
    
             board := Board {Size: 3,}
             b := board.Create() 
           
             p1 := Player{Symbol: "X", Type: "human",}    
             c, err := b.MakeMove(0, "X")
             
	     g.Assert(err).Equal(nil)
             g.Assert(Status(c, p1)).Equal("NO WINNER YET")   
         })

         g.It("game setup results into win for player 1", func() {
    
             board := Board {Size: 3,}
             b := board.Create() 
           
             p1 := Player{Symbol: "X", Type: "human",}    
             c, _ := b.MakeMove(0, "X")
             i, _ := c.MakeMove(1, "X")
             e, _ := i.MakeMove(4, "O")
             d, err := e.MakeMove(2, "X")
             
	     g.Assert(err).Equal(nil)
             g.Assert(Status(d, p1)).Equal("X (human) WINS")   
         })
        
         g.It("game setup results into win for player 2", func() {
    
             board := Board {Size: 3,}
             b := board.Create() 
           
             p2 := Player{Symbol: "O", Type: "ai",}
             
             c, _ := b.MakeMove(0, "O")
             i, _ := c.MakeMove(1, "O")
             e, _ := i.MakeMove(4, "X")
             d, err := e.MakeMove(2, "O")
             
             g.Assert(err).Equal(nil) 
             g.Assert(Status(d, p2)).Equal("O (ai) WINS")   
         })
         
        g.It("game setup results into tie game2", func() {
    
             board := Board {Size: 3,}
             b := board.Create() 
           
             p2 := Player{Symbol: "O", Type: "ai",}
                
             c, _ := b.MakeMove(0, "x")
             d, _ := c.MakeMove(1, "x")
             e, _ := d.MakeMove(2, "o")
             f, _ := e.MakeMove(3, "o")
             m, _ := f.MakeMove(4, "o")
             h, _ := m.MakeMove(5, "x")
             i, _ := h.MakeMove(6, "x")
             j, _ := i.MakeMove(7, "o")
             k, _ := j.MakeMove(8, "x")
             
             g.Assert(k.TieCombination()).Equal(true)
             g.Assert(Status(k, p2)).Equal("TIE GAME")
         })
      })
 
}
