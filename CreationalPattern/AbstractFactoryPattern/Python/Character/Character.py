
class Character:
    """
        Abstract class to define the character
    """
    def __init__(self, name):
        self.name = name

    def __str__(self):
        return self.name