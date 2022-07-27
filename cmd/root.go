/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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
	RunE:  runRoot,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&rootConfig.Input, "input", "i", "", "Help message for toggle")
	rootCmd.Flags().StringVarP(&rootConfig.Output, "output", "o", "", "Help message for toggle")
	rootCmd.Flags().BoolVarP(&rootConfig.Force, "force", "f", false, "Help message for toggle")
	rootCmd.Flags().BoolVar(&rootConfig.Indent, "indent", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&rootConfig.EscapeHTML, "escape-html", "e", false, "Help message for toggle")
	rootCmd.Flags().Uint32VarP(&rootConfig.OutputPermission, "output-permissions", "p", 0644, "Help message for toggle")
	err := rootCmd.MarkFlagRequired("input")
	if err != nil {
		panic(err)
	}
	err = rootCmd.MarkFlagRequired("output")
	if err != nil {
		panic(err)
	}
}

func runRoot(cmd *cobra.Command, args []string) error {
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
