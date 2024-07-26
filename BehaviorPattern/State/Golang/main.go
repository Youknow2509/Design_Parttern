package main

import "fmt"

// PhoneState interface
type PhoneState interface {
    PressButton(*PhoneContext)
}

// UnlockedState struct
type UnlockedState struct{}

func (u *UnlockedState) PressButton(context *PhoneContext) {
    fmt.Println("Button pressed: Executing various functions...")
}

// LockedState struct
type LockedState struct{}

func (l *LockedState) PressButton(context *PhoneContext) {
    fmt.Println("Button pressed: Showing unlock screen...")
}

// LowBatteryState struct
type LowBatteryState struct{}

func (lb *LowBatteryState) PressButton(context *PhoneContext) {
    fmt.Println("Button pressed: Showing charging screen...")
}

// PhoneContext struct
type PhoneContext struct {
    currentState PhoneState
}

func NewPhoneContext(initialState PhoneState) *PhoneContext {
    return &PhoneContext{currentState: initialState}
}

func (p *PhoneContext) SetState(state PhoneState) {
    p.currentState = state
}

func (p *PhoneContext) PressButton() {
    p.currentState.PressButton(p)
}

func main() {
    // Creating a phone with an initial state of LockedState
    phone := NewPhoneContext(&LockedState{})

    // Pressing button when the phone is locked
    phone.PressButton()

    // Changing state to Unlocked
    phone.SetState(&UnlockedState{})
    phone.PressButton()

    // Changing state to LowBattery
    phone.SetState(&LowBatteryState{})
    phone.PressButton()

    // Changing state back to Locked
    phone.SetState(&LockedState{})
    phone.PressButton()
}
