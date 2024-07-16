
package Plants;

import Interface.Plant;

/**
 * Class for Sunflower
 */
public class SunFlower implements Plant {
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

    @Override
    public void show() {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'show'");
    }
}