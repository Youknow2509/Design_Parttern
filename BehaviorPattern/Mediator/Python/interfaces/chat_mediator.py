from abc import ABC, abstractmethod

class ChatMediator(ABC):
    @abstractmethod
    def send_message(self, message, user):
        pass

    @abstractmethod
    def add_user(self, user):
        pass
