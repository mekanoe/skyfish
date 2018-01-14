package eventstream

import (
	"bytes"
	"encoding/json"
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
