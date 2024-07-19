package Concretes;

import Interfaces.GameObject;

public class Plant implements GameObject {
    private final String type; // Intrinsic state
    private final String image;

    public Plant(String type, String image) {
        this.type = type;
        this.image = image;
    }

    @Override
    public void display(int x, int y) {
        System.out.println("Plant: " + type + " with image: " + image + " is displayed at (" + x + ", " + y + ")");
    }
}
