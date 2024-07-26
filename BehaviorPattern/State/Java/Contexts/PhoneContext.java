
package Contexts;

import Interfaces.PhoneState;

public class PhoneContext {
    private PhoneState currentState;

    public PhoneContext(PhoneState initialState) {
        this.currentState = initialState;
    }

    public void setState(PhoneState state) {
        this.currentState = state;
    }

    public void pressButton() {
        currentState.pressButton(this);
    }
}
