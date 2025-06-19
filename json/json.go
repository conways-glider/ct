// Package json provides encoding and decoding functionality for JSON format.
package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/conways-glider/ct/config"
)

// Encode converts the input data to JSON format with optional formatting options.
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

// Decode reads JSON data from the configured input reader and returns the parsed data.
func Decode(config *config.Config) (interface{}, error) {
	var out interface{}
	err := json.NewDecoder(config.InputReader).Decode(&out)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %w", err)
	}
	return out, nil
}
