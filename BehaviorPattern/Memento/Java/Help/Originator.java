package Help;
/**
 * The Originator class creates a Memento containing a snapshot of its current state
 * and uses the Memento to restore its state.
 */
public class Originator {
    private String state;

    /**
     * Sets the state of the Originator.
     * @param state The state to set.
     */
    public void setState(String state) {
        this.state = state;
    }

    /**
     * Gets the current state of the Originator.
     * @return The current state.
     */
    public String getState() {
        return state;
    }

    /**
     * Saves the current state to a new Memento.
     * @return A Memento containing the current state.
     */
    public Memento saveStateToMemento() {
        return new Memento(state);
    }

    /**
     * Restores the state from a Memento.
     * @param memento The Memento from which to restore the state.
     */
    public void getStateFromMemento(Memento memento) {
        state = memento.getState();
    }
}
