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
	     }
	     b := newBoard.Create() 
	     g.Assert(len(b.Array)).Equal(9)
 	})

	g.It("should create board size of 4", func() {
	     newBoard := board.Board{
               Size: 4,
	     }
	    b := newBoard.Create()
	    g.Assert(len(b.Array)).Equal(16)
	})
	
       g.It("should create board size of 7", func(){
               newBoard := board.Board{
               Size: 7,
	     }
	     b := newBoard.Create() 
	     g.Assert(len(b.Array)).Equal(49)
 	})
    })


    g.Describe("Make move on a board", func() {
       	 g.It("get available spots to be 9: for 3x3 board", func(){
         
            newBoard := board.Board{
              Size: 3,
	    }
            b := newBoard.Create()
            g.Assert(b.AvailableSpots).Equal(9)
        })

	g.It("makes move on position i and checks for available spots on board", func(){
	  
	  newBoard := board.Board{
	     Size: 3,
	  }
          position := 5
	  symbol := "x"
          my_board := newBoard.Create()
	  b, _ := my_board.MakeMove(position, symbol)
	  g.Assert(b.AvailableSpots).Equal(8)
	})
        
	g.It("makes two moves on position i and j, then checks for available spots on board", func(){
	  
          newBoard := board.Board{
	     Size: 3,
	  }

          positionOne := 5
	  symbolOne := "x"
          my_board := newBoard.Create()
	  b, _ := my_board.MakeMove(positionOne, symbolOne)
	  g.Assert(b.AvailableSpots).Equal(8)
          
          positionTwo := 2
	  symbolTwo := "y"
          c, _ := b.MakeMove(positionTwo, symbolTwo)
          g.Assert(c.AvailableSpots).Equal(7)
	})
        
       g.It("makes move in position (10) outside board range", func() {
   
          newBoard := board.Board{
	     Size: 3,
	  }
          
          position := 10
	  symbol := "x"
          my_board := newBoard.Create()
          _, err := my_board.MakeMove(position, symbol)
          g.Assert(err.Error()).Equal("position is Out of Range")
      })

       g.It("makes move in position (-1) outside board range", func() {
   
          newBoard := board.Board{
	     Size: 3,
	  }
          
          position := -1
	  symbol := "x"
          my_board := newBoard.Create()
          _, err := my_board.MakeMove(position, symbol)
          g.Assert(err.Error()).Equal("position is Out of Range")
      })


       g.It("makes move in position (4), that already has value/symbol", func() {
   
          newBoard := board.Board{
	     Size: 3,
	  }
          
          my_board := newBoard.Create()
          
          positionOne := 4
	  symbolOne := "x"
          b, err := my_board.MakeMove(positionOne, symbolOne)
          g.Assert(err).Equal(nil)
          g.Assert(b.AvailableSpots).Equal(8)

          positionTwo := 4
	  symbolTwo := "x"
          c, err := b.MakeMove(positionTwo, symbolTwo)
          g.Assert(err.Error()).Equal("position already filled")
          g.Assert(c.AvailableSpots).Equal(8)
      })
   }) 

     g.Describe("Winning Combination on a Board", func() {
         g.It("1st [horizontal] winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(1, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(2, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.HorizontalWinningCombination()).Equal(true)
         })

         g.It("2nd horizontal winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(3, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(5, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.HorizontalWinningCombination()).Equal(true)
         })

         g.It("3rd horizontal winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(6, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(7, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(8, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)
           g.Assert(e.HorizontalWinningCombination()).Equal(true)
         })

         g.It("1st [vertical] winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(3, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(6, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.VerticalWinningCombination()).Equal(true)
         })

         g.It("2nd vertical winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(1, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(7, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.VerticalWinningCombination()).Equal(true)
         })

         g.It("3rd vertical winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(2, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(5, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(8, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.VerticalWinningCombination()).Equal(true)
         })

         g.It("1st [diagonal] winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(8, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.DiagonalWinningCombination()).Equal(true)
         })

         g.It("2nd diagonal winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(2, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(6, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.DiagonalWinningCombination()).Equal(true)
         })
     })

     g.Describe("Check winning combination for different board states", func() {
         g.It(" 3 spots filled with no winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(1, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(4, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

           g.Assert(e.WinningCombination()).Equal(false)
         })

         g.It(" 4 spots filled with no winning combination", func(){
        
           newBoard := board.Board{
              Size: 3,
	   }
           b := newBoard.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(c.AvailableSpots).Equal(8)
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(1, "x")
           g.Assert(d.AvailableSpots).Equal(7)
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(4, "x")
           g.Assert(e.AvailableSpots).Equal(6)
           g.Assert(err).Equal(nil)

	   f, err := e.MakeMove(2, "y")
           g.Assert(f.AvailableSpots).Equal(5)
           g.Assert(err).Equal(nil)

           g.Assert(e.WinningCombination()).Equal(false)
         })
   })
}








