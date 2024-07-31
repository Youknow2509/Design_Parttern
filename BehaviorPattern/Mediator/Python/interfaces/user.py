from abc import ABC, abstractmethod

class User(ABC):
    def __init__(self, mediator, name):
        self.mediator = mediator
        self.name = name

    @abstractmethod
    def send(self, message):
        pass

    @abstractmethod
    def receive(self, message):
        pass
