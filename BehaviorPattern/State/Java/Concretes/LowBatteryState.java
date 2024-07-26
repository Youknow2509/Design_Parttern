package Concretes;

import Interfaces.PhoneState;
import Contexts.PhoneContext;

public class LowBatteryState implements PhoneState {
    @Override
    public void pressButton(PhoneContext context) {
        System.out.println("Button pressed: Showing charging screen...");
    }
}