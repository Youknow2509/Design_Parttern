
from Interfaces import PhoneState

class UnlockedState(PhoneState):
    def press_button(self, context):
        print("Button pressed: Executing various functions...")