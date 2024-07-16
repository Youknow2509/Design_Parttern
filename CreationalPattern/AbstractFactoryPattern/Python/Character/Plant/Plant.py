
from ..Character import Character

class Plant(Character):
    """
        Class to define the plant
    """
    def __init__(self, *arg, **kwargs):
        self.name = kwargs.get("name", "Name Plant")
        self.health = kwargs.get("health", 100)
        if arg:
            lenArg = len(arg)
            if lenArg >= 1:
                self.name = arg[0] 
            if lenArg >= 2:
                self.health = arg[1]  
        super().__init__(self.name)

    def __str__(self):
        return f"Plant: {self.name} - Health: {self.health}"