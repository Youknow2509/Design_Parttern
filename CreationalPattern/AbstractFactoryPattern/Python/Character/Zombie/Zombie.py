
from ..Character import Character

class Zombie(Character):
    """
        Class to define the zombie
    """
    def __init__(self, *arg, **kwargs):
        self.name = kwargs.get("name", "Name Zombie")
        self.health = kwargs.get("health", 100)
        self.dame = kwargs.get("dame", 10)

        if arg:
            lenArg = len(arg)
            if lenArg >= 1:
                self.name = arg[0] 
            if lenArg >= 2:
                self.health = arg[1]  
            if lenArg >= 3:
                self.dame = arg[2]
        super().__init__(self.name)

    def __str__(self):
        return f"Zombie: {self.name} - Health: {self.health}"
