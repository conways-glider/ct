package convert

import (
	"os"
	"strings"
	"testing"

	"github.com/conways-glider/ct/config"
	"github.com/stretchr/testify/assert"
)

func getTestConfig(inputFile string, output string, indent bool) *config.Config {
	config := config.Config{
		Input:  inputFile,
		Output: output,
		Indent: indent,
	}
	err := config.Validate()
	if err != nil {
		panic(err)
	}
	return &config
}

func TestConvert(t *testing.T) {
	type args struct {
		rootConfig *config.Config
	}
	tests := []struct {
		name     string
		args     args
		wantFile string
		wantErr  bool
	}{
		{
			name: "toml to json",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/toml.toml", "json", false),
			},
			wantFile: "../test_resources/files/json.json",
			wantErr:  false,
		},
		{
			name: "toml to json with indent",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/toml.toml", "json", true),
			},
			wantFile: "../test_resources/files/json_indent.json",
			wantErr:  false,
		},
		{
			name: "yaml to json",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/yaml.yaml", "json", false),
			},
			wantFile: "../test_resources/files/json.json",
			wantErr:  false,
		},
		{
			name: "yaml to json with indent",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/yaml.yaml", "json", true),
			},
			wantFile: "../test_resources/files/json_indent.json",
			wantErr:  false,
		},
		{
			name: "yml to json",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/yml.yml", "json", false),
			},
			wantFile: "../test_resources/files/json.json",
			wantErr:  false,
		},
		{
			name: "yml to json with indent",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/yml.yml", "json", true),
			},
			wantFile: "../test_resources/files/json_indent.json",
			wantErr:  false,
		},
		{
			name: "json to yaml",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/json.json", "yaml", false),
			},
			wantFile: "../test_resources/files/yaml.yaml",
			wantErr:  false,
		},
		{
			name: "yaml to toml",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/yaml.yaml", "toml", false),
			},
			wantFile: "../test_resources/files/toml.toml",
			wantErr:  false,
		},
		{
			name: "toml to hcl",
			args: args{
				rootConfig: getTestConfig("../test_resources/files/toml.toml", "hcl", false),
			},
			wantFile: "../test_resources/files/hcl.hcl",
			wantErr:  false,
		},
		// Removed due to bug in HCL resulting in things being converted to arrays
		// {
		// 	name: "hcl to yaml",
		// 	args: args{
		// 		rootConfig: getTestConfig("../test_resources/files/yaml.yaml", "hcl", false),
		// 	},
		// 	wantFile: "../test_resources/files/hcl.hcl",
		// 	wantErr:  false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			want, err := os.ReadFile(tt.wantFile)
			assert.Nil(err)

			got, err := Convert(tt.args.rootConfig)
			assert.Nil(err)

			// Remove newlines from the output for cross system testing
			wantedString := strings.ReplaceAll(string(want), "\n", "")
			wantedString = strings.ReplaceAll(wantedString, "\r", "")
			gotString := strings.ReplaceAll(string(got), "\n", "")
			gotString = strings.ReplaceAll(gotString, "\r", "")

			assert.Equal(wantedString, gotString)
		})
	}
}
