
class Light:
    """
        Class Light object that will be used as a receiver in the command pattern
    """
    def __init__(self):
        pass
    
    def on(self):
        print("Light is on")

    def off(self):
        print("Light is off")