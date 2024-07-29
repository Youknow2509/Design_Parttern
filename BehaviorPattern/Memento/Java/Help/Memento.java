package Help;

/**
 * The Memento class stores the state of the Originator.
 */
public class Memento {
    private final String state;

    /**
     * Constructor to set the state.
     * @param state The state to save in the Memento.
     */
    public Memento(String state) {
        this.state = state;
    }

    /**
     * Gets the saved state.
     * @return The state stored in the Memento.
     */
    public String getState() {
        return state;
    }
}

