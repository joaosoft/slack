package slack

import (
	"encoding/json"
	"time"
)

type Attachment struct {
	title      string `json:"title"`
	attachment string `json:"attachment"`
}

type Message struct {
	title       string        `json:"title"`
	action      string        `json:"action"`
	environment string        `json:"environment"`
	message     string        `json:"message"`
	code        string        `json:"code"`
	attachments []*Attachment `json:"attachments"`
	error       error         `json:"error"`
	occurredAt  *time.Time    `json:"occurred_at"`
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) Title(title string) *Message {
	m.title = title
	return m
}

func (m *Message) Action(action string) *Message {
	m.action = action
	return m
}

func (m *Message) Environment(environment string) *Message {
	m.environment = environment
	return m
}

func (m *Message) Message(message string) *Message {
	m.message = message
	return m
}

func (m *Message) Code(code string) *Message {
	m.code = code
	return m
}

func (m *Message) SetOccurredAt(time *time.Time) *Message {
	m.occurredAt = time
	return m
}

func (m *Message) Attachment(title string, attachment interface{}) *Message {
	textAttachment, _ := json.MarshalIndent(attachment, "", "\t")
	m.attachments = append(m.attachments, &Attachment{
		title:      title,
		attachment: string(textAttachment),
	})
	return m
}

func (m *Message) Error(err error) *Message {
	m.error = err
	return m
}
