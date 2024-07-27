package Concretes;

public class SupportRequest {
    private String message;
    private int complexity;

    public SupportRequest(String message, int complexity) {
        this.message = message;
        this.complexity = complexity;
    }

    public String getMessage() {
        return message;
    }

    public int getComplexity() {
        return complexity;
    }
}
