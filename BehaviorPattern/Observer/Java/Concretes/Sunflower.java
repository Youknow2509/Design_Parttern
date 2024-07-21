package Concretes;

import Interfaces.Observer;

public class Sunflower implements Observer {
    private String plantAction;

    @Override
    public void update(String action) {
        this.plantAction = action;
        System.out.println("Sunflower: Zombie action detected: " + action);
    }
}
