/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: State Pattern
 */

import Concretes.*;
import Interfaces.*;
import Contexts.PhoneContext;

public class main {
    public static void main(String[] args) {
        // Creating a phone with an initial state of LockedState
        PhoneContext phone = new PhoneContext(new LockedState());

        // Pressing button when the phone is locked
        phone.pressButton();

        // Changing state to Unlocked
        phone.setState(new UnlockedState());
        phone.pressButton();

        // Changing state to LowBattery
        phone.setState(new LowBatteryState());
        phone.pressButton();

        // Changing state back to Locked
        phone.setState(new LockedState());
        phone.pressButton();
    }
}