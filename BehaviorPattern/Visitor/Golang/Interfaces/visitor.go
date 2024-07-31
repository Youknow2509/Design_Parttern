package Interfaces

// Visitor interface defines visit methods for each concrete element
type Visitor interface {
    Visit(element Element)
}
