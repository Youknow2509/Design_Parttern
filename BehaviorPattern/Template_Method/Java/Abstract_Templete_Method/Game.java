package Abstract_Templete_Method;

public abstract class Game {
     // Template method
    public final void play() {
        initialize();
        startPlay();
        endPlay();
    }

    // Abstract methods to be implemented by subclasses
    abstract void initialize();
    abstract void startPlay();
    abstract void endPlay();
}
