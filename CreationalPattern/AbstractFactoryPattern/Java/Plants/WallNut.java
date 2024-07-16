
package Plants;

import Interface.Plant;

/**
 * Class for Wallnut
 */
public class WallNut implements Plant {
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

