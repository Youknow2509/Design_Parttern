package concretes

import "mediator_pattern/interfaces"

// ChatMediatorImpl is a concrete implementation of the ChatMediator interface.
type ChatMediatorImpl struct {
	users []interfaces.User
}

// NewChatMediatorImpl creates a new ChatMediatorImpl.
func NewChatMediatorImpl() *ChatMediatorImpl {
	return &ChatMediatorImpl{}
}

// SendMessage sends a message to all users except the sender.
func (c *ChatMediatorImpl) SendMessage(message string, user interfaces.User) {
	for _, u := range c.users {
		if u != user {
			u.Receive(message)
		}
	}
}

// AddUser adds a user to the mediator.
func (c *ChatMediatorImpl) AddUser(user interfaces.User) {
	c.users = append(c.users, user)
}
