package cmd

import (
	"os"

	"github.com/conways-glider/ct/config"
	"github.com/conways-glider/ct/convert"
	"github.com/conways-glider/ct/output"
	"github.com/spf13/cobra"
)

var rootConfig *config.Config = &config.Config{}

var rootCmd = &cobra.Command{
	Use:   "ct",
	Short: "Config Transformer",
	Long: `Config Tranfromer (ct) is a tool to convert between YAML, TOML, JSON, and HCL

The Input and Output flags take either a file or extension.

If input is a file, the file will be read, inferring the type from the file extension.
If input is a type (e.g. toml, yaml, json, or hcl), it will be read from stdin.

If output is a file, the file will be written, inferring the type from the file extension.
If output is a type (e.g. toml, yaml, json, or hcl), it will be written to stdout.

It supports the following formats: TOML, YAML, JSON, and HCL.`,
	RunE: runRoot,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&rootConfig.Input, "input", "i", "", "Input file or extension (e.g. example.toml or toml) (accepted extensions: toml, yaml, json, hcl)")
	rootCmd.Flags().StringVarP(&rootConfig.Output, "output", "o", "", "Output file or extension  (e.g. example.json or json) (accepted extensions: toml, yaml, json, hcl)")
	rootCmd.Flags().BoolVarP(&rootConfig.Force, "force", "f", false, "Force overwrite of output file")
	rootCmd.Flags().BoolVar(&rootConfig.Indent, "indent", false, "Indent output (JSON & TOML only)")
	rootCmd.Flags().BoolVarP(&rootConfig.EscapeHTML, "escape-html", "e", false, "Escapes HTML (JSON only)")
	rootCmd.Flags().Uint32VarP(&rootConfig.OutputPermission, "output-permissions", "p", config.DefaultFilePermission, "File permissions for output file")
	err := rootCmd.MarkFlagRequired("input")
	if err != nil {
		panic(err)
	}
	err = rootCmd.MarkFlagRequired("output")
	if err != nil {
		panic(err)
	}
}

func runRoot(_ *cobra.Command, _ []string) error {
	if err := rootConfig.Validate(); err != nil {
		return err
	}

	convertedData, err := convert.Convert(rootConfig)
	if err != nil {
		return err
	}

	err = output.Output(rootConfig, convertedData)
	if err != nil {
		return err
	}

	return nil
}
