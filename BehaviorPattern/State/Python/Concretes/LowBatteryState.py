
from Interfaces import PhoneState

class LowBatteryState(PhoneState):
    def press_button(self, context):
        print("Button pressed: Showing charging screen...")