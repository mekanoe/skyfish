package eventstream

import (
	"fmt"

	"github.com/fatih/color"
)

func (es *EventStream) Progress(cur, max float64) *Event {
	return es.Event("progress",
		fmt.Sprintf("%f out of %f", cur, max),
	).WithData(
		"max", max,
	).WithData(
		"current", cur,
	).WithColor(
		color.New(color.FgGreen),
	)
}

func (es *EventStream) State(state string) *Event {
	return es.Event("state", state).WithColor(color.New(color.FgHiRed))
}
