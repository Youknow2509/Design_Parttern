
package Concretes;

import Interfaces.PhoneState;
import Contexts.PhoneContext;

public class UnlockedState implements PhoneState {
    @Override
    public void pressButton(PhoneContext context) {
        System.out.println("Button pressed: Executing various functions...");
    }
}