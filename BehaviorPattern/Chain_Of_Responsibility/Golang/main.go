package main

import (
    "v/Concretes"
)

func main() {
    level1Handler := &concretes.Level1SupportHandler{}
    level2Handler := &concretes.Level2SupportHandler{}
    level3Handler := &concretes.Level3SupportHandler{}

    level1Handler.SetNext(level2Handler)
    level2Handler.SetNext(level3Handler)

    request1 := NewSupportRequest("Simple issue", 1)
    request2 := NewSupportRequest("Moderate issue", 2)
    request3 := NewSupportRequest("Complex issue", 3)
    request4 := NewSupportRequest("Very complex issue", 4)

    level1Handler.HandleRequest(request1)
    level1Handler.HandleRequest(request2)
    level1Handler.HandleRequest(request3)
    level1Handler.HandleRequest(request4)
}
