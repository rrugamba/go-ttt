package player

type Player struct {
   Symbol, Type string
}

func (player Player) GetSymbol() string {
  return player.Symbol
}

func (player Player) GetType() string {
  return player.Type
}


