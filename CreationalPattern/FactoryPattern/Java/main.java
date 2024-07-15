/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Factory Pattern
 */

enum TypePlant {
    SUNFLOWER,
    PEASHOOTER,
    WALLNUT,
    // TODO enums
}

interface Plant {
    void show();
}

class SunFlower implements Plant {
    // Var parameters
    private String name;
    private int health;
    private int price;
    // TODO parameters

    // Constructor
    public SunFlower() {
        super();
        this.name = "SunFlower";
        this.health = 100;
        this.price = 50;
    }

    public SunFlower(String name, int health, int price) {
        super();
        this.name = name;
        this.health = health;
        this.price = price;
    }

    // Override method show
    @Override
    public void show() {
        System.out.println("Name: " + name);
        System.out.println("Health: " + health);
        System.out.println("Price: " + price);
    }

    // Getter and Setter 
    // TODO add ...
}

class PeaShooter implements Plant {
    // Var parameters
    private String name;
    private int health;
    private int dame;
    private int price;
    // TODO parameters

    // Constructor
    public PeaShooter() {
        super();
        this.name = "PeaShooter";
        this.health = 100;
        this.dame = 10;
        this.price = 100;
    }

    public PeaShooter(String name, int health, int dame, int price) {
        super();
        this.name = name;
        this.health = health;
        this.price = price;
        this.dame = dame;
    }

    // Override method show
    @Override
    public void show() {
        System.out.println("Name: " + name);
        System.out.println("Health: " + health);
        System.out.println("Price: " + price);
        System.out.println("Dame: " + dame);
    }

    // Getter and Setter 
    // TODO add ...
}

class WallNut implements Plant {
    // Var parameters
    private String name;
    private int health;
    private int price;
    // TODO parameters

    // Constructor
    public WallNut() {
        super();
        this.name = "WallNut";
        this.health = 200;
        this.price = 50;
    }

    public WallNut(String name, int health, int price) {
        super();
        this.name = name;
        this.health = health;
        this.price = price;
    }

    // Override method show
    @Override
    public void show() {
        System.out.println("Name: " + name);
        System.out.println("Health: " + health);
        System.out.println("Price: " + price);
    }

    // Getter and Setter 
    // TODO add ...
}

class PlantFactory {
    public static Plant getPlant(TypePlant type) {
        switch (type) {
            case SUNFLOWER:
                return new SunFlower();
            case PEASHOOTER:
                return new PeaShooter();
            case WALLNUT:
                return new WallNut();
            default:
                System.out.println("Don't have type plant: " + type);
                return null;
        }
    }
}

public class main {
    public static void main(String[] args) {
        Plant sunFlower = PlantFactory.getPlant(TypePlant.SUNFLOWER);
        sunFlower.show();
        System.out.println("====================================");

        Plant peaShooter = PlantFactory.getPlant(TypePlant.PEASHOOTER);
        peaShooter.show();
        System.out.println("====================================");


        Plant wallNut = PlantFactory.getPlant(TypePlant.WALLNUT);
        wallNut.show();
        System.out.println("====================================");

        return;
    }
}