
package receivers

import (
	"fmt"
)

/**
 * Class Light object that will be used as a receiver in the command pattern
 */        
type Light struct {
	
}

// Light on 
func (l *Light) On() {
	fmt.Println("Light is on")
}

// Light off 
func (l *Light) Off() {
	fmt.Println("Light is off")
}