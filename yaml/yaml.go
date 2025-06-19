// Package yaml provides encoding and decoding functionality for YAML format.
package yaml

import (
	"bytes"
	"fmt"

	"github.com/conways-glider/ct/config"
	"gopkg.in/yaml.v3"
)

// Encode converts the input data to YAML format.
func Encode(in interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := yaml.NewEncoder(buf).Encode(in)
	if err != nil {
		return nil, fmt.Errorf("could not encode yaml: %w", err)
	}

	return buf.Bytes(), nil
}

// Decode reads YAML data from the configured input reader and returns the parsed data.
func Decode(config *config.Config) (interface{}, error) {
	var out interface{}
	err := yaml.NewDecoder(config.InputReader).Decode(&out)
	if err != nil {
		return nil, fmt.Errorf("could not decode yaml: %w", err)
	}

	return out, nil
}
