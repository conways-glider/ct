package yaml

import (
	"bytes"
	"fmt"

	"github.com/conways-glider/ct/config"
	"gopkg.in/yaml.v3"
)

func Encode(in interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := yaml.NewEncoder(buf).Encode(in)
	if err != nil {
		return nil, fmt.Errorf("could not encode toml: %w", err)
	}

	return buf.Bytes(), nil
}

func Decode(config *config.Config) (interface{}, error) {
	var out interface{}
	err := yaml.NewDecoder(config.InputReader).Decode(&out)
	if err != nil {
		return nil, fmt.Errorf("could not decode toml: %w", err)
	}

	return out, nil
}
