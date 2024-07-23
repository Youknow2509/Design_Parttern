package concretes

import (
	"fmt"
	"v/Abstract_Template_Methods"
)

type Cricket struct{}

// Constructor
func NewCricket() abstracttemplatemethods.Game {
	return &Cricket{}
}

func (c Cricket) Initialize() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}

func (c Cricket) StartPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func (c Cricket) EndPlay() {
	fmt.Println("Cricket Game Finished!")
}