package convert

import (
	"github.com/conways-glider/ct/config"
	"github.com/conways-glider/ct/json"
	"github.com/conways-glider/ct/toml"
	"github.com/conways-glider/ct/yaml"
)

func Convert(rootConfig *config.Config) ([]byte, error) {
	var decoded interface{}
	var err error = nil

	switch rootConfig.InputExtension {
	case config.JSON:
		decoded, err = json.Decode(rootConfig)
	case config.TOML:
		decoded, err = toml.Decode(rootConfig)
	case config.YAML:
		decoded, err = yaml.Decode(rootConfig)
	}

	if err != nil {
		return nil, err
	}

	var encoded []byte
	switch rootConfig.OutputExtension {
	case config.JSON:
		encoded, err = json.Encode(rootConfig, decoded)
	case config.TOML:
		encoded, err = toml.Encode(rootConfig, decoded)
	case config.YAML:
		encoded, err = yaml.Encode(rootConfig, decoded)
	}

	if err != nil {
		return nil, err
	}

	return encoded, nil
}
