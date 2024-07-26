
from Interfaces import PhoneState

class PhoneContext:
    def __init__(self, initial_state: PhoneState):
        self.current_state = initial_state

    def set_state(self, state: PhoneState):
        self.current_state = state

    def press_button(self):
        self.current_state.press_button(self)
