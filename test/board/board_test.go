package test_Board

import (
        "testing"
        . "github.com/franela/goblin"
        "../../src/board"
)

func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe ("Creating a Board of any size", func() {
	g.It("should create board size of 3", func(){
               newBoard := board.Board{
               Size: 3,
	       AvailableSpots: 9,
	     }
	     b := newBoard.Create() 
	     g.Assert(len(b)).Equal(9)
 	})

	g.It("should create a 4 size board", func() {
	     newBoard := board.Board{
               Size: 4,
	       AvailableSpots: 16,
	     }
	    b := newBoard.Create()
	    g.Assert(len(b)).Equal(16)
	})
    })

    g.Describe("Make move on a board", func() {
	g.It("get available spots to be 9: for 3x3 board", func(){
               newBoard := board.Board{
                 Size: 3,
	         AvailableSpots: 9,
	       }
            g.Assert(newBoard.AvailableSpots).Equal(9)
       })

	g.It("makes move on position i and checks for available spots on board", func(){
	  
	  newBoard := board.Board{
	     Size: 3,
	     AvailableSpots: 9,
	  }
          position := 5
	  symbol := "x"
	  b, _ := newBoard.MakeMove(position, symbol)
	  
	  g.Assert(b.AvailableSpots).Equal(8)
	})

        
	g.It("makes two moves on position i and j, then checks for available spots on board", func(){
	  
          newBoard := board.Board{
	     Size: 3,
	     AvailableSpots: 9,
	  }

          positionOne := 5
	  symbolOne := "x"
	  b, _ := newBoard.MakeMove(positionOne, symbolOne)
	  g.Assert(b.AvailableSpots).Equal(8)
          
          positionTwo := 2
	  symbolTwo := "y"
          c, _ := b.MakeMove(positionTwo, symbolTwo)
          g.Assert(c.AvailableSpots).Equal(7)
	})
        
       g.It("makes move in position outside board range", func() {
   
          newBoard := board.Board{
	     Size: 3,
	     AvailableSpots: 9,
	  }
          
          position := 10
	  symbol := "x"
          _, err := newBoard.MakeMove(position, symbol)
          g.Assert(err.Error()).Equal("position is Out of Range")
      })












    })
}
