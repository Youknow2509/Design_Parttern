
package concretes

import (
	"v/Receivers"
	"v/Interfaces"
)

/**
 * Class Light Off Command implementation Command Interface
 */
type LightOffCommand struct {
	Light receivers.Light
}

// Constructor
func NewLightOffCommand(light receivers.Light) interfaces.Command {
	return &LightOffCommand{Light: light}
}

//  Override Execute command
func (l *LightOffCommand) Execute() {
	l.Light.Off()
}