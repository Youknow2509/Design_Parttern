/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Abstract Factory Pattern
 */

import Enum.*;
import Interface.*;
import Plants.*;
import Zombies.*;

/**
 * Interface Factory
 */
interface InterfaceFactory {
    Plant plantFactory(PlantType plantType);
    Zombie zombieFactory(ZombieType zombieType);
}

/**
 * Deloy factory plant
 */
class PlantFactory implements InterfaceFactory {
    // Factory Plant
    @Override
    public Plant plantFactory(PlantType plantType) {
        switch (plantType) {
            case SUNFLOWER:
                return new SunFlower();
            case PEASHOOTER:
                return new PeaShooter();
            case WALLNUT:
                return new WallNut();
            default:
                System.out.println("Plant Type: " + plantType + " is not valid");
                return null;
        }
    }

    // Factory Zombie
    @Override 
    public Zombie zombieFactory(ZombieType zombieType) {
        System.out.println("Factory Plant don't create Zombie");
        return null;
    }
}

/** 
 * Deloy factory zombie
 */
class ZombieFactory implements InterfaceFactory {
    // Factory Plant
    @Override
    public Plant plantFactory(PlantType plantType) {
        System.out.println("Factory Zombie don't create Plant");
        return null;
    }

    // Factory Zombie
    @Override 
    public Zombie zombieFactory(ZombieType zombieType) {
        switch (zombieType) {
            case NORMAL:
                return new NormalZombie();
            case FLAG:
                return new FlagZombie();
            default:
                System.out.println("Zombie Type: " + zombieType + " is not valid");
                return null;
        }
    }
}

/**
 * Class AbstractFactoryProducer
 */
class FactoryProducer {
    // Get Type Factory
    public static InterfaceFactory getFactoryType(CharacterType characterType) {
        switch (characterType) {
            case PLANT:
                return new PlantFactory();
            case ZOMBIE:
                return new ZombieFactory();
            default:
                System.out.println("Character Type: " + characterType + " is not valid");
                return null;
        }
    
    }
}

public class main {
    public static void main(String[] args) {

        InterfaceFactory producer = FactoryProducer.getFactoryType(CharacterType.PLANT);
        Plant plant = producer.plantFactory(PlantType.SUNFLOWER);
        plant.display();

        producer = FactoryProducer.getFactoryType(CharacterType.ZOMBIE);
        Zombie zombie = producer.zombieFactory(ZombieType.NORMAL);
        zombie.display();

        return;
    }
}
