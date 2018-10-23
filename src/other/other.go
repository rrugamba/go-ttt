package other

import "../player"

type Players struct {
   Array []player.Player
}

func (players Players) SwitchPlayers(p player.Player) player.Player {
    if p.Type == players.Array[0].Type {
       return players.Array[1]
    }
    return players.Array[0]
}
