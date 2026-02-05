// Package downloader provides commands for downloading resources.
package downloader

import (
	"fmt"

	"github.com/spf13/cobra"
)

func retrieveFlagValues(cmd *cobra.Command) (*flagValues, error) {
	file, err := cmd.Flags().GetString(FlagsConfig.File.Name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving 'file' flag: %w", err)
	}

	outputDir, err := cmd.Flags().GetString(FlagsConfig.Output.Name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving 'output' flag: %w", err)
	}

	return &flagValues{filePath: file, outputDir: outputDir}, nil
}

// DownloadFromTxtFile entry point for downloading images from a text file.
func DownloadFromTxtFile(cmd *cobra.Command, args []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occurred: %v\n", r)
		}
	}()

	fmt.Println("Starting image download...")
	flags, err := retrieveFlagValues(cmd)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loading flag values %v", *flags)
}
