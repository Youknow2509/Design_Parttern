/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Memento Pattern
 */

package main

import (
	"fmt"
	"v/Helpers"
)

func main() {
	originator := &helpers.Originator{}
	caretaker := &helpers.Caretaker{}

	// Set initial states and save them
	originator.SetState("State #1")
	originator.SetState("State #2")
	caretaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #3")
	caretaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #4")

	// Display the current state
	fmt.Println("Current State:", originator.GetState())

	// Restore to previous states
	originator.GetStateFromMemento(caretaker.Get(0))
	fmt.Println("First saved State:", originator.GetState())

	originator.GetStateFromMemento(caretaker.Get(1))
	fmt.Println("Second saved State:", originator.GetState())
}
