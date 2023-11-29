package toml

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/conways-glider/ct/config"
)

func Encode(config *config.Config, in interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	encoder := toml.NewEncoder(buf)
	if !config.Indent {
		encoder.Indent = ""
	}
	err := encoder.Encode(in)
	if err != nil {
		return nil, fmt.Errorf("could not encode toml: %w", err)
	}

	return buf.Bytes(), nil
}

func Decode(config *config.Config) (interface{}, error) {
	var out interface{}

	_, err := toml.NewDecoder(config.InputReader).Decode(&out)
	if err != nil {
		return nil, fmt.Errorf("could not decode toml: %w", err)
	}

	return out, nil
}
