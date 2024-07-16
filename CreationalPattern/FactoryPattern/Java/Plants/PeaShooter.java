
package Plants;

import Interface.Plant;

public class PeaShooter implements Plant {
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

