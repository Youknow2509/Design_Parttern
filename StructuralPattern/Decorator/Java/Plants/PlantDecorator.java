package Plants;

import Interfaces.Plant;

public abstract class PlantDecorator implements Plant {
    // Variables
    private Plant plant;

    // Constructors
    public PlantDecorator(Plant plant) {
        this.plant = plant;
    }

    @Override
    public void attack() {
        plant.attack();
    }
}
