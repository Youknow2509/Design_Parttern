
from .Plant import Plant

class SunFlower(Plant):
    """
        Class to define the SunFlower
    """
    def __init__(self, *arg, **kwargs):
        self.name = kwargs.get("name", "SunFlower")
        self.health = kwargs.get("health", 100)
        self.sun = kwargs.get("sun", 10)

        if arg:
            lenArg = len(arg)
            if lenArg >= 1:
                self.name = arg[0] 
            if lenArg >= 2:
                self.health = arg[1]  
            if lenArg >= 3:
                self.sun = arg[2]
        super().__init__(*arg, **kwargs)

    def __str__(self):
        return f"SunFlower: {self.name} - Health: {self.health} -sun: {self.sun}"
    
    # TODO methods, params ...