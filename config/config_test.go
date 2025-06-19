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
				Input:  "json",
				Output: "yaml",
			},
			wantErr: false,
		},
		{
			name: "invalid input extension",
			config: Config{
				Input:  "invalid",
				Output: "yaml",
			},
			wantErr: true,
		},
		{
			name: "invalid output extension",
			config: Config{
				Input:  "json",
				Output: "invalid",
			},
			wantErr: true,
		},
		{
			name: "yml extension normalization",
			config: Config{
				Input:  "yml",
				Output: "yaml",
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
			if !tt.wantErr && config.Input == "yml" {
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
		{"config.json", "json"},
		{"data.YAML", "yaml"},
		{"file.toml", "toml"},
		{"test.HCL", "hcl"},
		{"nested/path/file.yml", "yml"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, getExtension(tt.fileName))
	}
}
