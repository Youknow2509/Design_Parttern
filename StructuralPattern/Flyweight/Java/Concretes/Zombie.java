package Concretes;

import Interfaces.GameObject;

public class Zombie implements GameObject {
    private final String type; // Intrinsic state
    private final String image;

    public Zombie(String type, String image) {
        this.type = type;
        this.image = image;
    }

    @Override
    public void display(int x, int y) {
        System.out.println("Zombie: " + type + " with image: " + image + " is displayed at (" + x + ", " + y + ")");
    }
}
