package Concretes;

import Interfaces.SupportHandler;

public class Level2SupportHandler extends SupportHandler {
    @Override
    public void handleRequest(SupportRequest request) {
        if (request.getComplexity() <= 2) {
            System.out.println("Level 2 support handled the request: " + request.getMessage());
        } else if (nextHandler != null) {
            nextHandler.handleRequest(request);
        }
    }
}
