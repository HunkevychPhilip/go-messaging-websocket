package types

import "encoding/json"

type Message struct {
	Owner   string
	Content string
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
