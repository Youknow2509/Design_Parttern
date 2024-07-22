
package invokers

import (
	"v/Interfaces"
)

/**
 * Class Remote Control
 */
type RemoteControl struct {
    command interfaces.Command
}

// Set command
func (r *RemoteControl) SetCommand(command interfaces.Command) {
	r.command = command
}

// Execute command
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}
