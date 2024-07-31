package concretes

import (
	"fmt"
	"mediator_pattern/interfaces"
)

// UserImpl is a concrete implementation of the User interface.
type UserImpl struct {
	mediator interfaces.ChatMediator
	name     string
}

// NewUserImpl creates a new UserImpl.
func NewUserImpl(mediator interfaces.ChatMediator, name string) *UserImpl {
	return &UserImpl{
		mediator: mediator,
		name:     name,
	}
}

// Send sends a message through the mediator.
func (u *UserImpl) Send(message string) {
	fmt.Printf("%s: Sending Message = %s\n", u.name, message)
	u.mediator.SendMessage(message, u)
}

// Receive receives a message.
func (u *UserImpl) Receive(message string) {
	fmt.Printf("%s: Received Message = %s\n", u.name, message)
}

// GetName returns the name of the user.
func (u *UserImpl) GetName() string {
	return u.name
}
