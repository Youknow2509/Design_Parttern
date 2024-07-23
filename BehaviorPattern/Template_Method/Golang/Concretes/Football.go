package concretes

import (
	"fmt"
	"v/Abstract_Template_Methods"
)

type Football struct{}

// Constructor
func NewFootball() abstracttemplatemethods.Game {
	return &Football{}
}

func (f Football) Initialize() {
	fmt.Println("Football Game Initialized! Start playing.")
}

func (f Football) StartPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (f Football) EndPlay() {
	fmt.Println("Football Game Finished!")
}