package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func New(config *Config) *Slack {
	return &Slack{
		config: config,
	}
}

func (s *Slack) Send(msg *Message) error {
	message := SlackMessage{
		Username: s.config.App,
		Text:     NewFormat(""),
	}

	if msg.title != "" {
		message.Text.Concat(Bold(NewFormat(msg.title)), NewLine())
	}

	if msg.action != "" {
		message.Text.Concat(BlockQuote(Bold(Text("Action:"))), Space(), NewFormat(msg.action), NewLine())
	}

	if msg.environment != "" {
		message.Text.Concat(BlockQuote(Bold(Text("Environment:"))), Space(), NewFormat(msg.environment), NewLine())
	}

	if msg.message != "" {
		message.Text.Concat(BlockQuote(Bold(Text("Message:"))), Space(), NewFormat(msg.message), NewLine())
	}

	if msg.error != nil {
		message.Text.Concat(BlockQuote(Bold(Text("Error:")).Concat(Space(), NewFormat(msg.error.Error()))))
	}

	if msg.code != "" {
		message.Text.Concat(BlockQuote(Bold(Text("Envelope:"))), NewLine(), Code(msg.code), NewLine())
	}

	if msg.occurredAt == nil {
		now := time.Now()
		msg.occurredAt = &now
	}

	message.Text.Concat(BlockQuote(Bold(Text("Occurred At:")), Space(), Text(msg.occurredAt.Format("2006-01-02 15:04:05"))))

	for _, att := range msg.attachments {
		message.Attachments = append(message.Attachments,
			&SlackAttachment{
				Title: att.title,
				Text:  att.attachment,
			})
	}

	payload, _ := json.Marshal(message)
	_, err := http.Post(s.config.Url, "application/json", bytes.NewBuffer(payload))

	return err
}

func NewFormat(text string) *Format {
	format := Format(text)
	return &format
}

func (f *Format) Concat(formats ...*Format) *Format {
	var join string
	for _, format := range formats {
		join += fmt.Sprintf("%s", format)
	}
	*f = *NewFormat(fmt.Sprintf("%s%s", f, join))

	return f
}

func Join(formats ...*Format) *Format {
	var join string
	for _, format := range formats {
		join += fmt.Sprintf("%s", format)
	}
	return NewFormat(join)
}

func applyFormat(option option, text *Format) *Format {
	return NewFormat(fmt.Sprintf(options[option], text))
}

func Quote(formats ...*Format) *Format {
	return applyFormat(optionQuote, Join(formats...))
}

func BlockQuote(formats ...*Format) *Format {
	return applyFormat(optionBlockQuote, Join(formats...))
}

func Bold(formats ...*Format) *Format {
	return applyFormat(optionBold, Join(formats...))
}

func Code(code string) *Format {
	return applyFormat(optionCode, NewFormat(code))
}

func Text(text string) *Format {
	return NewFormat(text)
}

func Space() *Format {
	return NewFormat(" ")
}

func NewLine() *Format {
	return applyFormat(optionNewLine, NewFormat(""))
}

func (f *Format) String() string {
	return string(*f)
}
