// Package downloader provides commands for downloading resources.
package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// DownloadFromTxtFile entry point for downloading images from a text file.
func DownloadFromTxtFile(cmd *cobra.Command, args []string) {
	startTime := time.Now()

	defer handlePanic()
	defer func() {
		fmt.Printf("Images downloading took %s\n", time.Since(startTime))
	}()

	flags, err := retrieveFlagValues(cmd)

	if err != nil {
		panic(err)
	}

	imgUrls := extractImageURLsFromFile(flags.filePath)
	createOutDir(flags.outputDir)

	wg := sync.WaitGroup{}
	resChan := make(chan string)

	for _, url := range imgUrls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			downloadImage(url, flags.outputDir, resChan)
		}(url)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for res := range resChan {
		fmt.Println(res)
	}
}

func downloadImage(imageUrl, outputDir string, resChan chan<- string) {
	startTime := time.Now()
	resChan <- fmt.Sprintf("Downloading image %s\n", imageUrl)

	filename := path.Base(imageUrl)
	destination := filepath.Join(outputDir, filename)
	file, err := os.Create(destination)

	if err != nil {
		resChan <- fmt.Sprintf("Error creating file for %s | %v\n", imageUrl, err)
		return
	}

	defer file.Close()

	res, err := http.Get(imageUrl)

	if err != nil {
		resChan <- fmt.Sprintf("Error downloading image %s | %v\n", imageUrl, err)
		return
	}

	defer res.Body.Close()

	if _, err = io.Copy(file, res.Body); err != nil {
		resChan <- fmt.Sprintf("Error saving image %s | %v\n", imageUrl, err)
		return
	}

	resChan <- fmt.Sprintf("Image %s downloaded successfully in %s of time!\n", imageUrl, time.Since(startTime))
}

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

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Printf("An error occurred: %v\n", r)
	}
}
