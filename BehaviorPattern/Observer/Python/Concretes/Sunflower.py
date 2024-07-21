
from Interfaces import Observer

class Sunflower(Observer):
    """
        Class Sunflower implements the interface Observer
    """
    def __init__(self):
        super().__init__()
        self.plantAction: str = None
        
    def update(self, action):
        self.plantAction = action
        print("Sunflower: Zombie action detected: " + action)
    

