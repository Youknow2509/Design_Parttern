package main

type SupportRequest struct {
    Message    string
    Complexity int
}

func NewSupportRequest(message string, complexity int) *SupportRequest {
    return &SupportRequest{
        Message:    message,
        Complexity: complexity,
    }
}

func (sr *SupportRequest) GetMessage() string {
    return sr.Message
}

func (sr *SupportRequest) GetComplexity() int {
    return sr.Complexity
}
