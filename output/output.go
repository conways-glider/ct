package output

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/conways-glider/ct/config"
)

// Output writes the converted data either to a file or to stdout
// based on the configuration settings.
func Output(rootConfig *config.Config, output []byte) error {
	if rootConfig.OutputIsFile {
		filemode := fs.FileMode(rootConfig.ParsedOutputPermission)
		err := os.WriteFile(rootConfig.Output, output, filemode)
		if err != nil {
			return fmt.Errorf("error writing output file: %w", err)
		}
	} else {
		_, err := fmt.Print(string(output))
		if err != nil {
			return fmt.Errorf("error writing to stdout: %w", err)
		}
	}
	return nil
}
