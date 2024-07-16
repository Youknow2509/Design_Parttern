/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Singleton Pattern
 */

class Person {

    private String name;
    private String contact;

    public Person(String name, String contact) {
        this.name = name;
        this.contact = contact;
    }

    public void display() {
        System.out.println("Name: " + name + " - Contact: " + contact);
    }
}

public class main {

    // Variables
    private static Person instance = null;

    private static Person getInstance() {
        if (instance == null) {
            instance = new Person("Ly Tran Vinh", "lytranvinh.work@gmail.com");
            System.out.println("CreatedPerson");
        } else {
            System.out.println("Person already created");
        }
        return instance;
    }

    public static void main(String[] args) {
        
        for (int i = 0; i < 3; i++) {
            Person person = getInstance();
            person.display();
        }

        return;
    }
}
