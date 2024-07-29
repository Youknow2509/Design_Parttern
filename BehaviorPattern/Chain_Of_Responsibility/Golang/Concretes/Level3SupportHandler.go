package concretes

import (
    "fmt"
    "v"
    "v/Interfaces"
)

type Level3SupportHandler struct {
    nextHandler interfaces.SupportHandler
}

func (h *Level3SupportHandler) SetNext(handler interfaces.SupportHandler) {
    h.nextHandler = handler
}

func (h *Level3SupportHandler) HandleRequest(request *v.SupportRequest) {
    if request.GetComplexity() <= 3 {
        fmt.Println("Level 3 support handled the request:", request.GetMessage())
    } else {
        fmt.Println("Request is too complex to be handled.")
    }
}
