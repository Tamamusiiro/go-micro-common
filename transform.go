package common

import "encoding/json"

func Transform(from, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, to)
}
