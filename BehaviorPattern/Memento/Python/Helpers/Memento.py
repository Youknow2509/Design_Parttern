class Memento:
    def __init__(self, state):
        """
        Initializes the Memento with the given state.
        :param state: The state to save in the Memento.
        """
        self._state = state

    def get_state(self):
        """
        Returns the saved state.
        :return: The state stored in the Memento.
        """
        return self._state
