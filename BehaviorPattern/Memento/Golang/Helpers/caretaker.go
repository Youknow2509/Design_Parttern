package helpers

// Caretaker is responsible for managing Mementos
type Caretaker struct {
	mementoList []*Memento
}

// Add adds a Memento to the list
func (c *Caretaker) Add(m *Memento) {
	c.mementoList = append(c.mementoList, m)
}

// Get retrieves a Memento from the list by index
func (c *Caretaker) Get(index int) *Memento {
	if index < 0 || index >= len(c.mementoList) {
		return nil
	}
	return c.mementoList[index]
}
