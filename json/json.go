package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/conways-glider/ct/config"
)

func Encode(config *config.Config, in interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(config.EscapeHTML)
	if config.Indent {
		encoder.SetIndent("", "  ")
	}
	err := encoder.Encode(in)
	if err != nil {
		return nil, fmt.Errorf("could not encode json: %w", err)
	}

	return buf.Bytes(), err
}

func Decode(config *config.Config) (interface{}, error) {
	var out interface{}
	err := json.NewDecoder(config.InputReader).Decode(&out)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %w", err)
	}
	return out, nil
}
