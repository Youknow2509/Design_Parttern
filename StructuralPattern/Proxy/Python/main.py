
from Interfaces import Image
from Concretes import RealImage

class ProxyImage(Image):
    """
        Class ProxyImage implementation for interface image
    """
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        # TODO ...
        if args:
            self.fileName = str(args[0])
        self.realImage = None

    def display(self): 
        if self.realImage is None:
            self.realImage = RealImage(self.fileName)

        self.realImage.display()

def main():
    """
        Main function
    """
    proxyImage = ProxyImage("test_10mb.jpg")
    # Image will be loaded from disk
    proxyImage.display()
    print("")
    # Image will not be loaded from disk
    proxyImage.display()
    
if __name__ == '__main__':
    main()