package json

import (
	"bytes"
	"encoding/json"
)

func FormatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "  ")
	if err != nil {
		return data, err
	}
	return out.Bytes(), nil
}
