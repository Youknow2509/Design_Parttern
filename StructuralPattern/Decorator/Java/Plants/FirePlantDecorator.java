package Plants;

import Interfaces.Plant;

public class FirePlantDecorator extends PlantDecorator {

    // Constructors
    public FirePlantDecorator(Plant decoratedPlant) {
        super(decoratedPlant);
    }

    @Override
    public void attack() {
        super.attack();
        System.out.print(" Adding fire attack.");
    }
}
