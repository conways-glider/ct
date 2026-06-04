package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid json to yaml",
			config: Config{
				Input:  JSON,
				Output: YAML,
			},
			wantErr: false,
		},
		{
			name: "invalid input extension",
			config: Config{
				Input:  "invalid",
				Output: YAML,
			},
			wantErr: true,
		},
		{
			name: "invalid output extension",
			config: Config{
				Input:  JSON,
				Output: "invalid",
			},
			wantErr: true,
		},
		{
			name: "yml extension normalization",
			config: Config{
				Input:  YML,
				Output: YAML,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			// Create a temporary config to avoid modifying the test case
			config := tt.config

			err := config.Validate()
			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}

			// Check yml normalization
			if !tt.wantErr && config.Input == YML {
				assert.Equal("yaml", config.InputExtension, "Expected input extension to be normalized to 'yaml'")
			}
		})
	}
}

func TestGetExtension(t *testing.T) {
	tests := []struct {
		fileName string
		want     string
	}{
		{"config.json", JSON},
		{"data.YAML", YAML},
		{"file.toml", TOML},
		{"test.HCL", HCL},
		{"nested/path/file.yml", YML},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, getExtension(tt.fileName))
	}
}
