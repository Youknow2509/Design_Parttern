from interfaces import ChatMediator

class ChatMediatorImpl(ChatMediator):
    def __init__(self):
        self.users = []

    def send_message(self, message, user):
        for u in self.users:
            if u != user:
                u.receive(message)

    def add_user(self, user):
        self.users.append(user)
