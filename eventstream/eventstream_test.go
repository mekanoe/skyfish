package eventstream

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func getES() (es *EventStream, buf *bytes.Buffer) {
	buf = &bytes.Buffer{}
	es = NewEventStream(buf, ESJson)
	return
}

func getData(b []byte) (map[string]interface{}, error) {
	var o map[string]interface{}

	err := json.Unmarshal(b, &o)
	if err != nil {
		return o, err
	}

	data := o["data"].(map[string]interface{})
	return data, err
}

func TestEvent(t *testing.T) {
	es, buf := getES()

	es.Event(map[string]string{"hello": "world nyaaa"})

	dat, err := getData(buf.Bytes())
	if err != nil {
		t.Error(err)
		return
	}

	if dat["hello"] != "world nyaaa" {
		t.Error("data differed")
	}
}

func TestDummyDev(t *testing.T) {
	es := NewEventStream(os.Stdout, ESDev)

	es.Event(map[string]string{"myaaa": "nyaaa"})

}
