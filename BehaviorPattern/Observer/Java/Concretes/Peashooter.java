package Concretes;

import Interfaces.Observer;

public class Peashooter implements Observer {
    private String plantAction;

    @Override
    public void update(String action) {
        this.plantAction = action;
        System.out.println("Peashooter: Zombie action detected: " + action);
    }
}
