
from Interfaces import PhoneState

class LockedState(PhoneState):
    def press_button(self, context):
        print("Button pressed: Showing unlock screen...")