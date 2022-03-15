package websocket

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

func parseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")

	return json.Unmarshal([]byte(data.String()), target)
}
