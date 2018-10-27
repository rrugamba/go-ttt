package board_test

import (
        "testing"
        . "github.com/franela/goblin"
        ."../../src/board"
)

func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe ("Creating a Board of any size", func() {
	g.It("should create board size of 3", func(){
               board := Board{
               Size: 3,
	     }
	     b := board.Create() 
	     g.Assert(len(b.Array)).Equal(9)
 	})

	g.It("should create board size of 4", func() {
	     board := Board{
               Size: 4,
	     }
	    b := board.Create()
	    g.Assert(len(b.Array)).Equal(16)
	})
	
       g.It("should create board size of 7", func(){
               board := Board{
               Size: 7,
	     }
	     b := board.Create() 
	     g.Assert(len(b.Array)).Equal(49)
 	})
    })


    g.Describe("Make move on a board", func() {
       	 g.It("get available spots to be 9: for 3x3 board", func(){
         
            board := Board{
              Size: 3,
	    }
            b := board.Create()
            g.Assert(len(b.GetAvailableSpots())).Equal(9)
        })

	g.It("makes move on position i and checks for available spots on board", func(){
	  
	  board := Board{
	     Size: 3,
	  }
          position := 5
	  symbol := "x"
          my_board := board.Create()
	  b, _ := my_board.MakeMove(position, symbol)
	  g.Assert(len(b.GetAvailableSpots())).Equal(8)
	})
        
	g.It("makes two moves on position i and j, then checks for available spots on board", func(){
	  
          board := Board{
	     Size: 3,
	  }

          positionOne := 5
	  symbolOne := "x"
          my_board := board.Create()
	  b, _ := my_board.MakeMove(positionOne, symbolOne)
	  g.Assert(len(b.GetAvailableSpots())).Equal(8)
          
          positionTwo := 2
	  symbolTwo := "y"
          c, _ := b.MakeMove(positionTwo, symbolTwo)
          g.Assert(len(c.GetAvailableSpots())).Equal(7)
	})
        
       g.It("makes move in position (10) outside board range", func() {
   
          board := Board{
	     Size: 3,
	  }
          
          position := 10
	  symbol := "x"
          my_board := board.Create()
          _, err := my_board.MakeMove(position, symbol)
          g.Assert(err.Error()).Equal("position is Out of Range")
      })

       g.It("makes move in position (-1) outside board range", func() {
   
          board := Board{
	     Size: 3,
	  }
          
          position := -1
	  symbol := "x"
          my_board := board.Create()
          _, err := my_board.MakeMove(position, symbol)
          g.Assert(err.Error()).Equal("position is Out of Range")
      })


       g.It("makes move in position (4), that already has value/symbol", func() {
   
          board := Board{
	     Size: 3,
	  }
          
          my_board := board.Create()
          
          positionOne := 4
	  symbolOne := "x"
          b, err := my_board.MakeMove(positionOne, symbolOne)
          g.Assert(err).Equal(nil)
          g.Assert(len(b.GetAvailableSpots())).Equal(8)

          positionTwo := 4
	  symbolTwo := "x"
          c, err := b.MakeMove(positionTwo, symbolTwo)
          g.Assert(err.Error()).Equal("position already filled")
          g.Assert(len(c.GetAvailableSpots())).Equal(8)
      })
   }) 

     g.Describe("Winning Combination on a Board", func() {
         g.It("1st [horizontal] winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(1, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(2, "x")
           g.Assert(err).Equal(nil)

           g.Assert(e.HorizontalWinningCombination()).Equal(true)
         })

         g.It("2nd horizontal winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(3, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(5, "x")
           g.Assert(err).Equal(nil)

           g.Assert(e.HorizontalWinningCombination()).Equal(true)
         })

         g.It("3rd horizontal winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(6, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(7, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(8, "x")
           g.Assert(err).Equal(nil)
           g.Assert(e.HorizontalWinningCombination()).Equal(true)
         })

         g.It("1st [vertical] winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(3, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(6, "x")
           g.Assert(err).Equal(nil)
           g.Assert(e.VerticalWinningCombination()).Equal(true)
         })

         g.It("2nd vertical winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(1, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(7, "x")
           g.Assert(err).Equal(nil)
           g.Assert(e.VerticalWinningCombination()).Equal(true)
         })

         g.It("3rd vertical winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(2, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(5, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(8, "x")
           g.Assert(err).Equal(nil)
           g.Assert(e.VerticalWinningCombination()).Equal(true)
         })

         g.It("1st [diagonal] winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(8, "x")
           g.Assert(err).Equal(nil)
           g.Assert(e.DiagonalWinningCombination()).Equal(true)
         })

         g.It("2nd diagonal winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(2, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(4, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(6, "x")
           g.Assert(err).Equal(nil)
           g.Assert(e.DiagonalWinningCombination()).Equal(true)
         })
     })

     g.Describe("Check for array of available spots", func() {
         g.It(" 3 spots filled, array should be of size 6", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, _ := b.MakeMove(0, "x")
	   d, _ := c.MakeMove(1, "x")
	   e, _ := d.MakeMove(4, "x")

           g.Assert(len(e.GetAvailableSpots())).Equal(6)
           g.Assert(e.GetAvailableSpots()).Equal([]int{2,3,5,6,7,8})
         })
         
         g.It(" 5 spots filled, array should be of size 4", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, _ := b.MakeMove(0, "x")
	   d, _ := c.MakeMove(1, "x")
	   e, _ := d.MakeMove(4, "x")
	   f, _ := e.MakeMove(3, "o")
	   h, _ := f.MakeMove(5, "o")

           g.Assert(len(h.GetAvailableSpots())).Equal(4)
           g.Assert(h.GetAvailableSpots()).Equal([]int{2,6,7,8})
         })
    })
     g.Describe("Check winning combination for different board states", func() {
         g.It(" 3 spots filled with no winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(1, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(4, "x")
           g.Assert(err).Equal(nil)

           g.Assert(e.WinningCombination()).Equal(false)
         })

         g.It(" 4 spots filled with no winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, err := b.MakeMove(0, "x")
           g.Assert(err).Equal(nil)
	  
	   d, err := c.MakeMove(1, "x")
           g.Assert(err).Equal(nil)
    
	   e, err := d.MakeMove(4, "x")
           g.Assert(err).Equal(nil)

	   f, err := e.MakeMove(2, "y")
           g.Assert(err).Equal(nil)
           g.Assert(f.WinningCombination()).Equal(false)
         })

         g.It(" set up a tie game", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
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
         })

         g.It(" set up a all board with no tie game and winning combination", func(){
        
           board := Board{
              Size: 3,
	   }
           b := board.Create()
           
	   c, _ := b.MakeMove(0, "x")
	   d, _ := c.MakeMove(1, "x")
	   e, _ := d.MakeMove(2, "o")
	   f, _ := e.MakeMove(3, "o")
	   m, _ := f.MakeMove(4, "x")
	   h, _ := m.MakeMove(5, "x")
	   i, _ := h.MakeMove(6, "o")
	   j, _ := i.MakeMove(7, "o")
	   k, _ := j.MakeMove(8, "x")
          
          g.Assert(k.TieCombination()).Equal(false)
          g.Assert(k.WinningCombination()).Equal(true)
         })
   })
}








