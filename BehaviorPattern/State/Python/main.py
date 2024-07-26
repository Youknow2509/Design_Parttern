"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: State Pattern
"""

from Concretes import *
from Contexts import *


def main():
    """
        Main function
    """
    # Creating a phone with an initial state of LockedState
    phone = PhoneContext(LockedState())

    # Pressing button when the phone is locked
    phone.press_button()

    # Changing state to Unlocked
    phone.set_state(UnlockedState())
    phone.press_button()

    # Changing state to LowBattery
    phone.set_state(LowBatteryState())
    phone.press_button()

    # Changing state back to Locked
    phone.set_state(LockedState())
    phone.press_button()
    
if __name__ == "__main__":
    main()