/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Prototype Pattern
 */

class Person {
    // Variables
    private String name;
    private int age;
    private String address;
    private String birthday;
    private String phoneNumber;

    // Constructors
    public Person() {
    }
    
    public Person(String name, int age, String address, String birthday, String phoneNumber) {
        this.name = name;
        this.age = age;
        this.address = address;
        this.birthday = birthday;
        this.phoneNumber = phoneNumber;
    }

    public void display() {
        System.out.println("Name: " + name + ", Age: " + age);
    }
    
    public Person clone() {
        return new Person(name, age, address, birthday, phoneNumber);
    }

    // Getters and setters
    public String getName() {
        return name;
    }
    public int getAge() {
        return age;
    }
    public String getAddress() {
        return address;
    }
    public String getBirthday() {
        return birthday;
    }
    public String getPhoneNumber() {
        return phoneNumber;
    }
    public void setName(String name) {
        this.name = name;
    }
    public void setAge(int age) {
        this.age = age;
    }
    public void setAddress(String address) {
        this.address = address;
    }
    public void setBirthday(String birthday) {
        this.birthday = birthday;
    }
    public void setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
    }

}

public class main {
    public static void main(String[] args) {
        Person person1 = new Person("John", 25, "New York", "01/01/1995", "1234567890");
        Person person2 = person1.clone();
        person2.setName("Alice");
        person2.setAge(22);
        person1.display();
        person2.display();
    }
}
