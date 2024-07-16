/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Adapter Pattern
 */


/**
 * Adapter Vietnamese to English
 */
interface VietnameseTarget {
    void send(String words);
}

class VietnameseAdaptee {
    public void send(String words) {
        System.out.println("Send: " + words);
    }
}

class VietnameseAdapter implements VietnameseTarget {
    private VietnameseAdaptee adaptee;

    public VietnameseAdapter(VietnameseAdaptee adaptee) {
        this.adaptee = adaptee;
    }

    @Override
    public void send(String words) {
        adaptee.send(words);
    }
}

/**
 * Adapter English to Vietnamese
 */
interface EnglishTarget {
    void send(String words);
}

class EnglishAdaptee {
    public void send(String words) {
        System.out.println("Send: " + words);
    }
}

class EnglishAdapter implements EnglishTarget {
    private EnglishAdaptee adaptee;

    public EnglishAdapter(EnglishAdaptee adaptee) {
        this.adaptee = adaptee;
    }

    @Override
    public void send(String words) {
        adaptee.send(words);
    }
}

// Client code
public class main {
    public static void main(String[] args) {
        VietnameseAdaptee vietnameseAdaptee = new VietnameseAdaptee();
        VietnameseTarget vietnameseAdapter = new VietnameseAdapter(vietnameseAdaptee);
        vietnameseAdapter.send("Xin chào");

        EnglishAdaptee englishAdaptee = new EnglishAdaptee();
        EnglishTarget englishAdapter = new EnglishAdapter(englishAdaptee);
        englishAdapter.send("Hello");
    }
}

