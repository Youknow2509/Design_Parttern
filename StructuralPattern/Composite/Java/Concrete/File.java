package Concrete;

import Interface.FileSystemComponent;

public class File implements Interface.FileSystemComponent {
    // Variables of the file
    private String name;
    private String type;
    private double size;
    // TODO: Add more variables

    public File(String name, String type, double size) {
        this.name = name;
        this.type = type;
        this.size = size;
    }

    @Override
    public void showDetails() {
        System.out.println("File: " + name + " - Type: " + type + " - Size: " + size);
    }

    // Getter and Setter
    // TODO: ....
}