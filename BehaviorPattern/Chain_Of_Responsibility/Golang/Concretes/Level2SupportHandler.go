package concretes

import (
    "fmt"
    "v"
    "v/Interfaces"
)

type Level2SupportHandler struct {
    nextHandler interfaces.SupportHandler
}

func (h *Level2SupportHandler) SetNext(handler interfaces.SupportHandler) {
    h.nextHandler = handler
}

func (h *Level2SupportHandler) HandleRequest(request *v.SupportRequest) {
    if request.GetComplexity() <= 2 {
        fmt.Println("Level 2 support handled the request:", request.GetMessage())
    } else if h.nextHandler != nil {
        h.nextHandler.HandleRequest(request)
    }
}
