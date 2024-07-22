
package concretes

import (
	"v/Interfaces"
)

// Class NameCollection implementation Aggregate Interface
type NameCollection struct {
	names []string 
}

// NewNameCollection creates a new NameCollection
func NewNameCollection() interfaces.Aggregate {
	return &NameCollection{
		names: []string{"Alice", "Bob", "Charlie", "Diana", "Elisa"},
	}
}

// Create Iterator for NameCollection
func (n *NameCollection) CreateIterator() interfaces.Iterator {
	return &NameIterator{
		names: n.names,
		index: 0,
	}
}

// Class NameIterator implementation Iterator Interface
type NameIterator struct {
	names []string
	index int
}
    
// Constructor for NameIterator
func NewNameIterator(names []string) interfaces.Iterator {
	return &NameIterator{
		names: names,
		index: 0,
	}
}

// Check have next element
func (n *NameIterator) HasNext() bool {
	return n.index < len(n.names)
}
            
// Get next element
func (n *NameIterator) Next() interface{} {
	if n.HasNext() {
		name := n.names[n.index]
		n.index++
		return name
	}
	return nil
}
       