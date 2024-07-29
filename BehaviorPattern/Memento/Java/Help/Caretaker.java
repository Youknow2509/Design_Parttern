package Help;

import java.util.ArrayList;
import java.util.List;

/**
 * The Caretaker class is responsible for keeping the Mementos.
 */
public class Caretaker {
    private List<Memento> mementoList = new ArrayList<>();

    /**
     * Adds a Memento to the list.
     * @param state The Memento to add.
     */
    public void add(Memento state) {
        mementoList.add(state);
    }

    /**
     * Gets a Memento from the list.
     * @param index The index of the Memento to retrieve.
     * @return The retrieved Memento.
     */
    public Memento get(int index) {
        return mementoList.get(index);
    }
}
