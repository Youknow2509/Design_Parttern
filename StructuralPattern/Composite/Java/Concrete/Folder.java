package Concrete;

import Interface.FileSystemComponent;
import java.util.ArrayList;
import java.util.List;

public class Folder implements FileSystemComponent {
    // Variables of the folder
    private String name;
    private List<FileSystemComponent> children = new ArrayList<FileSystemComponent>();
    // TODO: add more variables

    // Constructor
    public Folder(String name) {
        this.name = name;
    }

    // Add a child to the folder
    public void add(FileSystemComponent child) {
        children.add(child);
    }

    // Remove a child from the folder
    public void remove(FileSystemComponent child) {
        children.remove(child);
    }

    // Show details about the folder
    @Override
    public void showDetails() {
        System.out.println("Folder: " + name);
        for (FileSystemComponent child : children) {
            child.showDetails();
        }
    }

    // Getter and setter
    // TODO ...
}