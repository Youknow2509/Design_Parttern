package interfaces

// User defines the interface for sending and receiving messages.
type User interface {
	Send(message string)
	Receive(message string)
	GetName() string
}
