"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Mediator Pattern
"""

from concretes import ChatMediatorImpl
from concretes import UserImpl

def main():
    mediator = ChatMediatorImpl()

    user1 = UserImpl(mediator, "John")
    user2 = UserImpl(mediator, "Jane")
    user3 = UserImpl(mediator, "Doe")

    mediator.add_user(user1)
    mediator.add_user(user2)
    mediator.add_user(user3)

    user1.send("Hello, everyone!")

if __name__ == "__main__":
    main()
