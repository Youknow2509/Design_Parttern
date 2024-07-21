package concretes

import (
	"v/Interfaces"
)

type Zombie struct {
	Observers []interfaces.Observer
	Action string
}

// Contructor
func NewZombie() interfaces.Subject {
	return &Zombie{}
}

func (zombie *Zombie) AddObserver(observer interfaces.Observer) {
	zombie.Observers = append(zombie.Observers, observer)
}

func (zombie *Zombie) RemoveObserver(observer interfaces.Observer) {
	for i, o := range zombie.Observers {
		if o == observer {
			zombie.Observers = append(zombie.Observers[:i], zombie.Observers[i+1:]...)
			break
		}
	}
}

func (zombie *Zombie) NotifyObservers() {
	for _, observer := range zombie.Observers {
		observer.Update(zombie.Action)
	}
}

func (z *Zombie) SetAction(action string) {
	z.Action = action
	z.NotifyObservers()
}

func (z *Zombie) GetAction() string {
	return z.Action
}