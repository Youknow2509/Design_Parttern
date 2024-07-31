/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Mediator Pattern
 */
import Concretes.ChatMediatorImpl;
import Concretes.UserImpl;
import Interfaces.ChatMediator;
import Interfaces.User;

public class main {
    public static void main(String[] args) {
        ChatMediator mediator = new ChatMediatorImpl();

        User user1 = new UserImpl(mediator, "John");
        User user2 = new UserImpl(mediator, "Jane");
        User user3 = new UserImpl(mediator, "Doe");

        mediator.addUser(user1);
        mediator.addUser(user2);
        mediator.addUser(user3);

        user1.send("Hello, everyone!");
    }
}
