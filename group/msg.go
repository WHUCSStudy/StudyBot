package group

type Message struct {
	Type    string `json:"type"`
	Text    string `json:"text,omitempty"`
	Url     string `json:"url,omitempty"`
	Id      int64  `json:"id,omitempty"`
	Time    int64  `json:"time,omitempty"`
	Target  int64  `json:"target,omitempty"`
	Display string `json:"display,omitempty"`
}

type MessageChain struct {
	messageChain []Message
}

func NewMessageChain() *MessageChain {
	return &MessageChain{make([]Message, 0)}
}
func (p *MessageChain) BuildImg(url string) *MessageChain {
	p.messageChain = append(p.messageChain, Message{Type: "Image", Url: url})
	return p
}

func (p *MessageChain) BuildText(text string) *MessageChain {
	p.messageChain = append(p.messageChain, Message{Type: "Plain", Text: text})
	return p
}

func (p *MessageChain) BuildAt() *MessageChain {
	return p
}
