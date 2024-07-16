
from .Zombie import Zombie

class FlagZombie(Zombie):
    """
        Class to define the flag zombie
    """
    def __init__(self, *arg, **kwargs):
        self.name = kwargs.get("name", "Flag Zombie")
        self.health = kwargs.get("health", 200)
        self.dame = kwargs.get("dame", 5)
        
        if arg:
            lenArg = len(arg)
            if lenArg >= 1:
                self.name = arg[0] 
            if lenArg >= 2:
                self.health = arg[1]  
            if lenArg >= 3:
                self.dame = arg[2]
        super().__init__(*arg, **kwargs)
        
    def __str__(self):
        return f"Flag Zombie: {self.name} - Health: {self.health} - Dame: {self.dame}"
    
    # TODO: Add methods, parameters ...