from .Memento import Memento

class Originator:
    def __init__(self):
        """
        Initializes the Originator with an empty state.
        """
        self._state = ""

    def set_state(self, state):
        """
        Sets the state of the Originator.
        :param state: The state to set.
        """
        self._state = state

    def get_state(self):
        """
        Returns the current state of the Originator.
        :return: The current state.
        """
        return self._state

    def save_state_to_memento(self):
        """
        Saves the current state to a new Memento.
        :return: A Memento containing the current state.
        """
        return Memento(self._state)

    def get_state_from_memento(self, memento):
        """
        Restores the state from a Memento.
        :param memento: The Memento from which to restore the state.
        """
        self._state = memento.get_state()
