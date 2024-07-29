package helpers

// Memento stores the state of the Originator
type Memento struct {
	state string
}

// NewMemento creates a new Memento with the given state
func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

// GetState returns the state stored in the Memento
func (m *Memento) GetState() string {
	return m.state
}