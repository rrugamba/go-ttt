package player

const HUMAN = "human"
const AI = "ai"

type Player struct {
   Symbol, Type string
}

func (player Player) GetSymbol() string {
  return player.Symbol
}

func (player Player) GetType() string {
  return player.Type
}

func (player Player) SwitchPlayer(a, b Player) Player {
     if player.Type == a.Type {
       return b
     }
     return a
}
