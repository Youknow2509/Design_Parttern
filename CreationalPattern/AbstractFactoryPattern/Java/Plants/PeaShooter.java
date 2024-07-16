
package Plants;

import Interface.Plant;

/**
 * Class for Peashooter
 */
public class PeaShooter implements Plant {
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

    @Override
    public void show() {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'show'");
    }
}
