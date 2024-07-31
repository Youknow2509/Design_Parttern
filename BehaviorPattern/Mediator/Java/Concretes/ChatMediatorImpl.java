package Concretes;

import Interfaces.ChatMediator;
import Interfaces.User;

import java.util.ArrayList;
import java.util.List;

public class ChatMediatorImpl implements ChatMediator {
    private List<User> users = new ArrayList<>();

    @Override
    public void sendMessage(String message, User user) {
        for (User u : users) {
            // Message should not be sent to the user who sent it
            if (u != user) {
                u.receive(message);
            }
        }
    }

    @Override
    public void addUser(User user) {
        users.add(user);
    }
}
