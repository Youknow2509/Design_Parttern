class Caretaker:
    def __init__(self):
        """
        Initializes the Caretaker with an empty list of Mementos.
        """
        self._memento_list = []

    def add(self, memento):
        """
        Adds a Memento to the list.
        :param memento: The Memento to add.
        """
        self._memento_list.append(memento)

    def get(self, index):
        """
        Retrieves a Memento from the list.
        :param index: The index of the Memento to retrieve.
        :return: The retrieved Memento.
        """
        return self._memento_list[index]
