/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Abstract Factory Pattern
 */

import Enum.*;

/**
 * Interface for Character
 */
interface Character {
    void display();
}

/**
 * Interface for Plant
 */
interface Plant extends Character {
    int getHealth();
}

/**
 * Interface for Zombie
 */
interface Zombie extends Character {
    int getDamage();
}

/**
 * Class for Sunflower
 */
class SunFlower implements Plant {
    // Var params for Sunflower
    private String name;
    private int hp;
    private int price;
    
    // Constructor
    public SunFlower() {
        this.name = "Sunflower";
        this.hp = 100;
        this.price = 50;
    }

    public SunFlower(String name, int hp, int price) {
        this.name = name;
        this.hp = hp;
        this.price = price;
    }
    
    // Override method Display from Plant
    @Override
    public void display() {
        System.out.println("Name: " + name + ", Price: " + price + ", HP: " + hp);
    }

    // Override method GetHealth from Plant
    @Override
    public int getHealth() {
        return hp;
    }

    // Getter and setter
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public void setHp(int hp) {
        this.hp = hp;
    }
}

/**
 * Class for Peashooter
 */
class PeaShooter implements Plant {
    // Var params for Peashooter
    private String name;
    private int hp;
    private int damage;
    private int price;
    
    // Constructor
    public PeaShooter() {
        this.name = "Peashooter";
        this.hp = 100;
        this.damage = 20;
        this.price = 100;
    }

    public PeaShooter(String name, int hp, int damage, int price) {
        this.name = name;
        this.hp = hp;
        this.damage = damage;
        this.price = price;
    }
    
    // Override method Display from Plant
    @Override
    public void display() {
        System.out.println("Name: " + name + ", Price: " + price + ", HP: " + hp + ", Damage: " + damage);
    }

    // Override method GetHealth from Plant
    @Override
    public int getHealth() {
        return hp;
    }

    // Getter and setter
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public int getDamage() {
        return damage;
    }

    public void setDamage(int damage) {
        this.damage = damage;
    }

    public void setHp(int hp) {
        this.hp = hp;
    }
}

/**
 * Class for Wallnut
 */
class WallNut implements Plant {
    // Var params for Wallnut
    private String name;
    private int hp;
    private int price;
    
    // Constructor
    public WallNut() {
        this.name = "Wallnut";
        this.hp = 300;
        this.price = 50;
    }

    public WallNut(String name, int hp, int price) {
        this.name = name;
        this.hp = hp;
        this.price = price;
    }
    
    // Override method Display from Plant
    @Override
    public void display() {
        System.out.println("Name: " + name + ", Price: " + price + ", HP: " + hp);
    }

    // Override method GetHealth from Plant
    @Override
    public int getHealth() {
        return hp;
    }

    // Getter and setter
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public void setHp(int hp) {
        this.hp = hp;
    }
}

/**
 * Class for NormalZombie
 */
class NormalZombie implements Zombie {
    // Var params for NormalZombie
    private String name;
    private int hp;
    private int damage;
    private int price;
    
    // Constructor
    public NormalZombie() {
        this.name = "NormalZombie";
        this.hp = 100;
        this.damage = 20;
        this.price = 50;
    }

    public NormalZombie(String name, int hp, int damage, int price) {
        this.name = name;
        this.hp = hp;
        this.damage = damage;
        this.price = price;
    }
    
    // Override method Display from Zombie
    @Override
    public void display() {
        System.out.println("Name: " + name + ", Price: " + price + ", HP: " + hp + ", Damage: " + damage);
    }

    // Override method GetDamage from Zombie
    @Override
    public int getDamage() {
        return damage;
    }

    // Getter and setter
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public void setDamage(int damage) {
        this.damage = damage;
    }

    public int getHp() {
        return hp;
    }

    public void setHp(int hp) {
        this.hp = hp;
    }
}

/**
 * Class for FlagZombie
 */
class FlagZombie implements Zombie {
    // Var params for FlagZombie
    private String name;
    private int hp;
    private int damage;
    private int price;
    
    // Constructor
    public FlagZombie() {
        this.name = "FlagZombie";
        this.hp = 100;
        this.damage = 20;
        this.price = 50;
    }

    public FlagZombie(String name, int hp, int damage, int price) {
        this.name = name;
        this.hp = hp;
        this.damage = damage;
        this.price = price;
    }
    
    // Override method Display from Zombie
    @Override
    public void display() {
        System.out.println("Name: " + name + ", Price: " + price + ", HP: " + hp + ", Damage: " + damage);
    }

    // Override method GetDamage from Zombie
    @Override
    public int getDamage() {
        return damage;
    }

    // Getter and setter
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public void setDamage(int damage) {
        this.damage = damage;
    }

    public int getHp() {
        return hp;
    }

    public void setHp(int hp) {
        this.hp = hp;
    }
}

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
