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
var RecordDeleted = DomioMessage{1004, "Record deleted succesfully"}
var CardDeleted = DomioMessage{1005, "Card deleted succesfully"}
var UserDeleted = DomioMessage{1006, "User deleted succesfully"}