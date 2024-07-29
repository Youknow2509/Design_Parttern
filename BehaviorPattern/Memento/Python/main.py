"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Memento Pattern
"""

from Helpers import *;

def main():
    """
    The main function.
    """
    # Create an Originator and a Caretaker
    originator = Originator()
    caretaker = Caretaker()

    # Set some states and save them
    originator.set_state("State #1")
    originator.set_state("State #2")
    caretaker.add(originator.save_state_to_memento())

    originator.set_state("State #3")
    caretaker.add(originator.save_state_to_memento())

    originator.set_state("State #4")

    # Display the current state
    print("Current State:", originator.get_state())

    # Restore the state from the first saved memento
    originator.get_state_from_memento(caretaker.get(0))
    print("First saved State:", originator.get_state())

    # Restore the state from the second saved memento
    originator.get_state_from_memento(caretaker.get(1))
    print("Second saved State:", originator.get_state())
    
if __name__ == '__main__':
    main()