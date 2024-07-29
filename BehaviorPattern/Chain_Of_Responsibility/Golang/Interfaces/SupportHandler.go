package interfaces

import "v"

type SupportHandler interface {
    SetNext(handler SupportHandler)
    HandleRequest(request *v.SupportRequest)
}
