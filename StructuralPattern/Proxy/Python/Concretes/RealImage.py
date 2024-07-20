from Interfaces import Image


class RealImage(Image):
    """
        Class RealImage implementation for interface image
    """
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.fileName = kwargs.get('fileName', None)
        # TODO ...
        if args:
            self.fileName = str(args[0])
        

    def RealImage(self, fileName):
        self.fileName = fileName
        self.selfloadFromDisk()

    def loadFromDisk(self):
        print("Loading " + self.fileName)
    
    def display(self):
        print("Displaying " + self.fileName)

