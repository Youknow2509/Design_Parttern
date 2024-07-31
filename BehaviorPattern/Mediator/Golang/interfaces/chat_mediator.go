package interfaces

// ChatMediator defines the interface for sending messages and adding users.
type ChatMediator interface {
	SendMessage(message string, user User)
	AddUser(user User)
}
