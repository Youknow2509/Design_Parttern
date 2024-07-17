"""
    Define shape
"""

from Color import Color

class Shape:
    """
        Interface for shape
    """
    def __init__(self, color):
        self.color = color
    
    def __str__(self):
        pass

class Circle(Shape):
    """
        Circle implements shape
    """
    def __init__(self,color):
        super().__init__(color)

    def __str__(self):
        return "Circle" + self.color.draw()
    
class Square(Shape):
    """
        Square implements shape
    """
    def __init__(self,color):
        super().__init__(color)

    def __str__(self):
        return "Square" + self.color.draw()
    