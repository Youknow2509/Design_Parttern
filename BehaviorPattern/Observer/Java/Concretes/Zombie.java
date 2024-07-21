package Concretes;

import java.util.ArrayList;
import java.util.List;

import Interfaces.Subject;
import Interfaces.Observer;

public class Zombie implements Subject {
    private List<Observer> observers = new ArrayList<>();
    private String action;

    @Override
    public void addObserver(Observer observer) {
        observers.add(observer);
    }
    @Override
    public void removeObserver(Observer observer) {
        observers.remove(observer);
    }
    @Override
    public void notifyObservers() {
        for (Observer observer : observers) {
            observer.update(action);
        }
    }

    public void setAction(String action) {
        this.action = action;
        notifyObservers();
    }

    public String getAction() {
        return action;
    }
}
