package eventstream

import (
	"github.com/fatih/color"
)

// Event is a meta-struct for data actually sent in the stream
type Event struct {
	Type     string
	Message  string
	Data     map[string]interface{}
	stream   *EventStream
	DevColor *color.Color
}

// WithData adds data annotations to the event output
func (e *Event) WithData(name string, value interface{}) *Event {
	e.Data[name] = value
	return e
}

func (e *Event) WithColor(c *color.Color) *Event {
	e.DevColor = c
	return e
}

func (e *Event) Send() error {
	return e.stream.outputEvent(e)
}
