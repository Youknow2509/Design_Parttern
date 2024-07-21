
from Interfaces import Observer

class Peashooter(Observer):
    """
        Class Peashooter implements the interface Observer
    """
    def __init__(self):
        super().__init__()
        self.plantAction: str = None
        
    def update(self, action):
        self.plantAction = action
        print("Peashooter: Zombie action detected: " + action)
    

