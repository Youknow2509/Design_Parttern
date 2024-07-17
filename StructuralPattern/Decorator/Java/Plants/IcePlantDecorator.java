package Plants;

import Interfaces.Plant;

public class IcePlantDecorator extends PlantDecorator {
    
    // Constructors
    public IcePlantDecorator(Plant plantDecorator) { 
        super(plantDecorator);
    }

    @Override
    public void attack() {
        super.attack();
        System.out.print(" Adding ice attack.");
    }
}
