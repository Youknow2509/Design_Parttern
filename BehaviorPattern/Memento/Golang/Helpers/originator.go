package helpers

// Originator manages its state and can create or restore a Memento
type Originator struct {
	state string
}

// SetState sets the state of the Originator
func (o *Originator) SetState(state string) {
	o.state = state
}

// GetState returns the current state of the Originator
func (o *Originator) GetState() string {
	return o.state
}

// SaveStateToMemento saves the current state to a new Memento
func (o *Originator) SaveStateToMemento() *Memento {
	return NewMemento(o.state)
}

// GetStateFromMemento restores the state from a Memento
func (o *Originator) GetStateFromMemento(m *Memento) {
	o.state = m.GetState()
}
