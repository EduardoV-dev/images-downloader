// Package cmd defines the CLI commands for the application.
package cmd

import (
	"os"

	"github.com/EduardoV-dev/images-downloader/internal/downloader"
	"github.com/EduardoV-dev/images-downloader/internal/flags"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "idd",
	Short: "Image downloader based on a images.txt file",
	Long: `Image downloader is a CLI tool that reads a text file containing image URLs
and downloads each image to a specified directory.`,
	Run: downloader.DownloadFromTxtFile,
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
	rootCmd.PersistentFlags().StringP(flags.FlagsConfig.File.Name, flags.FlagsConfig.File.Shorthand, flags.FlagsConfig.File.DefaultValue, flags.FlagsConfig.File.Usage)
	rootCmd.PersistentFlags().StringP(flags.FlagsConfig.Output.Name, flags.FlagsConfig.Output.Shorthand, flags.FlagsConfig.Output.DefaultValue, flags.FlagsConfig.Output.Usage)
	rootCmd.PersistentFlags().UintP(flags.FlagsConfig.Timeout.Name, flags.FlagsConfig.Timeout.Shorthand, flags.FlagsConfig.Timeout.DefaultValue, flags.FlagsConfig.Timeout.Usage)
}
