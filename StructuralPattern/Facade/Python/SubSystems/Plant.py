
class Plant:
    """
        Susclass plant
    """
    def __init__(self):
        self.name = 'Plant'

    def grow(self):
        print(self.name + " is growing")

    def attackZombie(self):
        print(self.name + " attack zombie")