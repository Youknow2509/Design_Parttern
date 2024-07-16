
from .Zombie import Zombie

class NormalZombie(Zombie):
    """
        Class to define the normal zombie
    """
    def __init__(self, *arg, **kwargs):
        self.name = kwargs.get("name", "Normal Zombie")
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
        super().__init__(*arg, **kwargs)
        
    def __str__(self):
        return f"Normal Zombie: {self.name} - Health: {self.health} - Dame: {self.dame}"
    
    # TODO: add methods, parameters ...