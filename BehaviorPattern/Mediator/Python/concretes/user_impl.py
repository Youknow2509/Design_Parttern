from interfaces import User

class UserImpl(User):
    def __init__(self, mediator, name):
        super().__init__(mediator, name)

    def send(self, message):
        print(f"{self.name}: Sending Message = {message}")
        self.mediator.send_message(message, self)

    def receive(self, message):
        print(f"{self.name}: Received Message = {message}")
