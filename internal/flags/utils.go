package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func GetFlagValues(cmd *cobra.Command) (*flagValues, error) {
	file, err := cmd.Flags().GetString(FlagsConfig.File.Name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving 'file' flag: %w", err)
	}

	if !strings.Contains(file, ".txt") {
		return nil, fmt.Errorf("File type not supported, make sure to use .txt file")
	}

	outputDir, err := cmd.Flags().GetString(FlagsConfig.Output.Name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving 'output' flag: %w", err)
	}

	timeout, err := cmd.Flags().GetUint(FlagsConfig.Timeout.Name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving 'timeout' flag: %w", err)
	}

	return &flagValues{FilePath: file, OutputDir: outputDir, Timeout: timeout}, nil
}
