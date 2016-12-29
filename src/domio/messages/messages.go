package messages

type DomioMessage struct {
    Code    int `json:"code"`
    Message string  `json:"message"`
}

func (e *DomioMessage) Description() string {
    return e.Message
}

var UserCreated = DomioMessage{1001, "User created successfully"}