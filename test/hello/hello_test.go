package test_hello

import (
	"testing"
	. "github.com/franela/goblin"
	 "../../src/hello"
)

func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe("Numbers", func() {
        
	g.It("Should add two numbers ", func() {
            g.Assert(hello.Add(1, 1)).Equal(2)
        })
        g.It("Should substract two numbers", func() {
            g.Assert(hello.Subtract(5,4)).Equal(1)
        })
        
	g.It("Should multiply  two numbers", func() {
          g.Assert(hello.Multiply(3, 4)).Equal(12)
	})
        
	g.It("Should divide two numbers ", func() {
            g.Assert(hello.Divide(3,1)).Equal(3)
        })
    })
}
