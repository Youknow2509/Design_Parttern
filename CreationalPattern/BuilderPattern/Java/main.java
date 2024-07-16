
class Person {
    // Variables
    private String firstName;
    private String lastName;
    private int age;
    private String address;
    private String phone;
    private String email;
    private String birthDate;
    private String favoriteSport;
    private String filmSport;
    
    // Constructors
    public Person() {
    }

    public Person(String firstName, String lastName, int age, String address, String phone, String email, String birthDate, String favoriteSport, String filmSport) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.age = age;
        this.address = address;
        this.phone = phone;
        this.email = email;
        this.birthDate = birthDate;
        this.favoriteSport = favoriteSport;
        this.filmSport = filmSport;
    }

    // Setters create builder pattern
    public Person setFirstName(String firstName) {
        this.firstName = firstName;
        return this;
    }

    public Person setLastName(String lastName) {
        this.lastName = lastName;
        return this;
    }

    public Person setAge(int age) {
        this.age = age;
        return this;
    }

    public Person setAddress(String address) {
        this.address = address;
        return this;
    }

    public Person setPhone(String phone) {
        this.phone = phone;
        return this;
    }

    public Person setEmail(String email) {
        this.email = email;
        return this;
    }

    public Person setBirthDate(String birthDate) {
        this.birthDate = birthDate;
        return this;
    }

    public Person setFavoriteSport(String favoriteSport) {
        this.favoriteSport = favoriteSport;
        return this;
    }

    public Person setFilmSport(String filmSport) {
        this.filmSport = filmSport;
        return this;
    }

    public Person build() {
        return new Person(firstName, lastName, age, address, phone, email, birthDate, favoriteSport, filmSport);
    }

    // To String
    @Override
    public String toString() {
        return "Person{" +
                "firstName='" + firstName + '\'' +
                ", lastName='" + lastName + '\'' +
                ", age=" + age +
                ", address='" + address + '\'' +
                ", phone='" + phone + '\'' +
                ", email='" + email + '\'' +
                ", birthDate='" + birthDate + '\'' +
                ", favoriteSport='" + favoriteSport + '\'' +
                ", filmSport='" + filmSport + '\'' +
                '}';
    }

    // Getters and setters
    // TODO add ...
}

public class main {
    public static void main(String[] args) {

        Person p = new Person()
            .setFirstName("John")
            .setLastName("Doe")
            .setAge(30)
            .setAddress("1234 Main St")
            .setPhone("123-456-7890")
            .build();

        System.out.println(p);

        Person p2 = new Person()
            .setFirstName("Tom")
            .setLastName("Mark")
            .setAge(25)
            .setAddress("1234 Main St")
            .setPhone("123-456-7890")
            .setEmail("tommark@gmail.com")
            .setBirthDate("01/01/1995")
            .setFavoriteSport("Soccer")
            .setFilmSport("Basketball")
            .build();

        System.out.println(p2);
    }
}
