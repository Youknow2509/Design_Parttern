/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Chain Of Responsibility Pattern
 */

import Concretes.SupportRequest;
import Concretes.Level1SupportHandler;
import Concretes.Level2SupportHandler;
import Concretes.Level3SupportHandler;
import Interfaces.SupportHandler;

public class main {
    public static void main(String[] args) {
        SupportHandler level1Handler = new Level1SupportHandler();
        SupportHandler level2Handler = new Level2SupportHandler();
        SupportHandler level3Handler = new Level3SupportHandler();

        level1Handler.setNextHandler(level2Handler);
        level2Handler.setNextHandler(level3Handler);

        SupportRequest request1 = new SupportRequest("Simple issue", 1);
        SupportRequest request2 = new SupportRequest("Moderate issue", 2);
        SupportRequest request3 = new SupportRequest("Complex issue", 3);
        SupportRequest request4 = new SupportRequest("Very complex issue", 4);

        level1Handler.handleRequest(request1);
        level1Handler.handleRequest(request2);
        level1Handler.handleRequest(request3);
        level1Handler.handleRequest(request4);
    }
}
