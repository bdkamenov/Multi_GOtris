package core

import "time"

type GameEntity struct {
	player1 Player
	player2 Player
	timeStarted time.Time
}

func (g *GameEntity) updatePlayers()  {

}