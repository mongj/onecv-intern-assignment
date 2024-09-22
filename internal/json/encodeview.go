package json

import (
	"encoding/json"
)

func EncodeView(view interface{}) ([]byte, error) {
	data, err := json.Marshal(view)
	if err != nil {
		return nil, err
	}
	return data, nil
}
