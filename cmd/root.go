// Package cmd defines the CLI commands for the application.
package cmd

import (
	"os"

	"github.com/EduardoV-dev/images-downloader/internal/downloader"
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
	rootCmd.PersistentFlags().StringP(downloader.FlagsConfig.File.Name, downloader.FlagsConfig.File.Shorthand, downloader.FlagsConfig.File.DefaultValue, downloader.FlagsConfig.File.Usage)
	rootCmd.PersistentFlags().StringP(downloader.FlagsConfig.Output.Name, downloader.FlagsConfig.Output.Shorthand, downloader.FlagsConfig.Output.DefaultValue, downloader.FlagsConfig.Output.Usage)
}
