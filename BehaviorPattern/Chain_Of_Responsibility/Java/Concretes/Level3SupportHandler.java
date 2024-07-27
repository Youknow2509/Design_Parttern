package Concretes;

import Interfaces.SupportHandler;

public class Level3SupportHandler extends SupportHandler {
    @Override
    public void handleRequest(SupportRequest request) {
        if (request.getComplexity() <= 3) {
            System.out.println("Level 3 support handled the request: " + request.getMessage());
        } else {
            System.out.println("Request is too complex to be handled.");
        }
    }
}
