"""
    Definitions for Color
"""

class Color:
    """
        Interface for Color objects
    """
    def __init__(self):
        pass

    # Define method: Draw
    def draw(self) -> str:
        pass

class RedColor(Color):
    """
        Deloy Color object - Red color
    """
    def __init__(self):
        pass
    
    # Deloy method: Draw
    def draw(self) -> str:
        return "Red Color"
    
class GreenColor(Color):
    """
        Deloy Color object - GreenColor
    """
    def __init__(self):
        pass
    
    # Deloy method: Draw
    def draw(self) -> str:
        return "Green Color"
    
class BlueColor(Color):
    """
        Deloy Color object - BlueColor
    """
    def __init__(self):
        pass
    
    # Deloy method: Draw
    def draw(self) -> str:
        return "Blue Color"