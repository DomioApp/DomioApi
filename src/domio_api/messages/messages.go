package messages

type DomioMessage struct {
    Code    int `json:"code"`
    Message string  `json:"message"`
}

func (e *DomioMessage) Description() string {
    return e.Message
}

var UserCreated = DomioMessage{1001, "User created successfully"}
var DomainDeleted = DomioMessage{1002, "Domain deleted succesfully"}
var SubscriptionDeleted = DomioMessage{1003, "Subscription deleted succesfully"}