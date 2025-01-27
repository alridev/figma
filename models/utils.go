package models

import (
	"bytes"
	"encoding/json"
)

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

func MapToStruct(input interface{}, output interface{}) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return
	}
	json.Unmarshal(jsonData, output)
}
