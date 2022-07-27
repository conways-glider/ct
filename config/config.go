package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/conways-glider/ct/slices"
)

type Config struct {
	// CLI flags
	Input            string
	Output           string
	Force            bool
	Indent           bool
	EscapeHTML       bool
	OutputPermission uint32

	// Internal values
	InputIsFile     bool
	OutputIsFile    bool
	InputExtension  string
	OutputExtension string
	InputReader     io.Reader
}

const fileRegex = `(?i)^.+\.(json|toml|yaml|yml)$`

const JSON = "json"
const TOML = "toml"
const YAML = "yaml"
const YML = "yml"

var validExtensions = []string{JSON, TOML, YAML, YML}

// Validate checks to see if the config is valid after flags are loaded and returns an error if not
func (config *Config) Validate() error {
	var returnedError error = nil
	re := regexp.MustCompile(fileRegex)

	// Check input string
	switch {
	case re.Match([]byte(config.Input)):
		config.InputIsFile = true
		config.InputExtension = getExtension(config.Input)
	case slices.Contains(validExtensions, strings.ToLower(config.Input)):
		config.Input = strings.ToLower(config.Input)
		config.InputExtension = config.Input
		config.InputIsFile = false
	default:
		errorString := fmt.Sprintf("invalid input type: %s", config.Input)
		returnedError = appendError(returnedError, errorString)
	}

	// Check output string
	switch {
	case re.Match([]byte(config.Output)):
		config.OutputIsFile = true
		config.OutputExtension = getExtension(config.Output)
	case slices.Contains(validExtensions, strings.ToLower(config.Output)):
		config.Output = strings.ToLower(config.Output)
		config.OutputExtension = config.Output
		config.OutputIsFile = false
	default:
		errorString := fmt.Sprintf("invalid output type: %s", config.Output)
		returnedError = appendError(returnedError, errorString)
	}

	// Handle yml => yaml
	if config.Input == YML {
		config.Input = YAML
	}
	if config.InputExtension == YML {
		config.InputExtension = YAML
	}
	if config.Output == YML {
		config.Output = YAML
	}
	if config.OutputExtension == YML {
		config.OutputExtension = YAML
	}

	// Check input file
	// If file does not exist, return error
	if config.InputIsFile {
		if _, err := os.Stat(config.Input); errors.Is(err, os.ErrNotExist) {
			errorString := fmt.Sprintf("input file does not exist: %s", config.Input)
			returnedError = appendError(returnedError, errorString)
		}
	}

	// Check output file
	// If file exists, return error
	if config.OutputIsFile && !config.Force {
		if _, err := os.Stat(config.Output); err == nil {
			errorString := fmt.Sprintf("output file already exists: %s, -f or changing the file name may help", config.Output)
			returnedError = appendError(returnedError, errorString)
		}
	}
	if returnedError != nil {
		return returnedError
	}

	if err := config.configureInput(); err != nil {
		return err
	}

	return nil
}

func (config *Config) configureInput() error {
	if config.InputIsFile {
		// Read file
		f, err := os.Open(config.Input)
		if err != nil {
			err = fmt.Errorf("error reading input file: %w", err)
			return err
		}

		config.InputReader = bufio.NewReader(f)
	} else {
		config.InputReader = os.Stdin
	}
	return nil
}

func getExtension(fileName string) string {
	fileExtension := filepath.Ext(fileName)
	fileExtension = strings.TrimPrefix(strings.ToLower(fileExtension), ".")
	return fileExtension
}

func appendError(err error, errorString string) error {
	if err == nil {
		return fmt.Errorf(errorString)
	}
	return fmt.Errorf("%w; %s", err, errorString)
}
