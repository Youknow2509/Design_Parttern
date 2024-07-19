/**
 * Struct game engine
 */
package subsystems

import (
	"fmt"
)

type GameEngine struct {}

// Constructor
func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

func (ge GameEngine) Start() {
	fmt.Println("Game engine started")
}

func (ge GameEngine) Stop() {
	fmt.Println("Game engine stopped")
}