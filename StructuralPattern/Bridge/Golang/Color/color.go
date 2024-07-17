package color

/**
 * Interface for color
 */
type Color interface {
	GetColor() string
}

/**
 * Red color implementation for color
 */
type RedColor struct {
}

// Deloy getColor method
func (c *RedColor) GetColor() string {
	return "Red"
}

/** 
 * Blue color implementation for color
 */
type BlueColor struct {
}

// Deloy getColor method
func (c *BlueColor) GetColor() string {
	return "Blue"
}