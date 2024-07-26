from abc import ABC, abstractmethod

class PhoneState(ABC):
    @abstractmethod
    def press_button(self, context):
        pass
