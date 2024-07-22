/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Command Pattern
 */

package main

import (
	"v/Concretes"
	"v/Receivers"
	"v/Invokers"
)

func main() {

    light := receivers.Light{}

    lightOn := concretes.NewLightOnCommand(light)
    lightOff := concretes.NewLightOffCommand(light)

    remote := invokers.RemoteControl{}

    // Turn the light on
    remote.SetCommand(lightOn)
    remote.PressButton()

    // Turn the light off
    remote.SetCommand(lightOff)
    remote.PressButton()
    
}