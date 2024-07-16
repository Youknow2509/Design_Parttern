
package Zombies;

import Interface.Zombie;

/**
 * Class for FlagZombie
 */
public class FlagZombie implements Zombie {
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
