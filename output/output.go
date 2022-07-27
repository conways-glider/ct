package output

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/conways-glider/ct/config"
)

func Output(rootConfig *config.Config, output []byte) error {
	if rootConfig.OutputIsFile {
		err := os.WriteFile(rootConfig.Output, output, fs.FileMode(rootConfig.OutputPermission))
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
