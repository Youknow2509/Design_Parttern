
package concretes

import (
	"v/Receivers"
	"v/Interfaces"
)

/**
 * Class Light On Command implementation Command Interface
 */
type LightOnCommand struct {
	Light receivers.Light
}

// Constructor
func NewLightOnCommand(light receivers.Light) interfaces.Command {
	return &LightOnCommand{Light: light}
}

//  Override Execute command
func (l *LightOnCommand) Execute() {
	l.Light.Off()
}