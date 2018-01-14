package eventstream

import (
	"fmt"
	"testing"
)

func TestEventCreate(t *testing.T) {
	es, buf := getES()
	es.Mode = ESDev

	ev := es.Event("myaa", "test helloooy nyaaa~")
	ev.WithData("nya?", true).Send()

	fmt.Println(buf)
}

func TestUsualEvents(t *testing.T) {
	es, buf := getES()
	es.Mode = ESDev

	es.Progress(50, 100).Send()
	es.State("test").Send()

	fmt.Println(buf)
}
