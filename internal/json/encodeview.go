package json

import (
	"encoding/json"
)

func EncodeView(view interface{}) ([]byte, error) {
	return json.Marshal(view)
}
