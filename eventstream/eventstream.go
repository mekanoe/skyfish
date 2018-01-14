package eventstream

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/fatih/color"
	"github.com/segmentio/ksuid"
)

const (
	ESDev  = 0
	ESJson = 1
)

var (
	colorBold = color.New(color.Bold)
)

type EventStream struct {
	Mode   int
	Output io.Writer

	jsonEnc *json.Encoder
}

func NewEventStream(o io.Writer, mode int) *EventStream {
	return &EventStream{
		Mode:    mode,
		Output:  o,
		jsonEnc: json.NewEncoder(o),
	}
}

func (es *EventStream) jsonEvent(e *Event) error {
	evt := map[string]interface{}{
		"timestamp": time.Now(),
		"id":        ksuid.New(),
		"data":      e.Data,
		"message":   e.Message,
		"type":      e.Type,
	}
	return es.jsonEnc.Encode(evt)
}

func (es *EventStream) devEvent(e *Event) error {
	typeName := e.Type
	if e.DevColor != nil {
		typeName = e.DevColor.Sprint(e.Type)
	}

	_, err := fmt.Fprintf(es.Output,
		"%s || %s\nMSG  > %s\nDATA > %v\n\n",
		colorBold.Sprint(`⚡️ EVENT`),
		typeName,
		e.Message,
		e.Data,
	)
	return err
}

func (es *EventStream) outputEvent(e *Event) error {
	switch es.Mode {
	case ESDev:
		return es.devEvent(e)
	default:
		return es.jsonEvent(e)
	}
}

func (es *EventStream) Event(typeName, msg string) *Event {
	return &Event{
		stream:  es,
		Message: msg,
		Type:    typeName,
		Data:    map[string]interface{}{},
	}
}
