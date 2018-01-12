package eventstream

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/segmentio/ksuid"
)

const (
	ESDev  = 0
	ESJson = 1
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

func (es *EventStream) jsonEvent(i interface{}) error {
	evt := map[string]interface{}{
		"timestamp": time.Now(),
		"id":        ksuid.New(),
		"data":      i,
	}
	return es.jsonEnc.Encode(evt)
}

func (es *EventStream) devEvent(i interface{}) error {
	_, err := fmt.Fprintf(es.Output, `⚡️ EVENT || %v\n`, i)
	return err
}

func (es *EventStream) Event(i interface{}) error {
	switch es.Mode {
	case ESDev:
		return es.devEvent(i)
	default:
		return es.jsonEvent(i)
	}
}
