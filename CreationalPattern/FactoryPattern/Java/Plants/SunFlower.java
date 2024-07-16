
package Plants;

import Interface.Plant;

public class SunFlower implements Plant {
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


