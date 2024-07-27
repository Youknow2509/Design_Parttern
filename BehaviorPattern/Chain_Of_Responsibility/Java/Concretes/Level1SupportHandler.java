package Concretes;

import Interfaces.SupportHandler;

public class Level1SupportHandler extends SupportHandler {
    @Override
    public void handleRequest(SupportRequest request) {
        if (request.getComplexity() <= 1) {
            System.out.println("Level 1 support handled the request: " + request.getMessage());
        } else if (nextHandler != null) {
            nextHandler.handleRequest(request);
        }
    }
}
