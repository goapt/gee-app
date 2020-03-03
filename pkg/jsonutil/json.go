package jsonutil

import (
	"encoding/json"
)

func JsonEncode(v interface{}) string {
	str, _ := json.Marshal(v)
	return string(str)
}
