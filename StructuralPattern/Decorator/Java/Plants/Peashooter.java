package Plants;

import Interfaces.Plant;

// Concrete interface Plant
public class Peashooter implements Plant {
    @Override
    public void attack() {
        System.out.print("Peashooter attacks.");
    }
}


