package abstracttemplatemethods

import (

)

// Game interface defines the methods that need to be implemented by concrete classes
type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
}

// play is the template method that defines the steps of the algorithm
func Play(g Game) {
	g.Initialize()
	g.StartPlay()
	g.EndPlay()
}