/**
 * Structor game facade
 */
package facade

import (
	"v/SubSystems"
)

type GameFacade struct {
	plant subsystems.Plant
	zombie subsystems.Zombie
	gameEngine subsystems.GameEngine
}

// Construct a new GameFacade
func NewGameFacade() *GameFacade {
	return &GameFacade{
		plant: *subsystems.NewPlant(),
		zombie: *subsystems.NewZombie(),
		gameEngine: *subsystems.NewGameEngine(),
	}
}

// Start the game
func (gf GameFacade) StartGame() {
	gf.gameEngine.Start()
	gf.plant.Grow()
	gf.zombie.Walk()
	gf.plant.Attack()
	gf.zombie.Attack()
	gf.gameEngine.Stop()
}

