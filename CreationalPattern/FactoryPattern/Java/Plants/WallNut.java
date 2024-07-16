
package Plants;

import Interface.Plant;

public class WallNut implements Plant {
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
