package concretes

import (
    "fmt"
    "v"
    "v/Interfaces"
)

type Level1SupportHandler struct {
    nextHandler interfaces.SupportHandler
}

func (h *Level1SupportHandler) SetNext(handler interfaces.SupportHandler) {
    h.nextHandler = handler
}

func (h *Level1SupportHandler) HandleRequest(request *v.SupportRequest) {
    if request.GetComplexity() <= 1 {
        fmt.Println("Level 1 support handled the request:", request.GetMessage())
    } else if h.nextHandler != nil {
        h.nextHandler.HandleRequest(request)
    }
}
