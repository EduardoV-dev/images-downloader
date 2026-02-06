// Package downloader provides commands for downloading resources.
package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/EduardoV-dev/images-downloader/internal/flags"
	"github.com/spf13/cobra"
)

var (
	fileLinesCount        = 0
	downloadedImagesCount = 0
	imagesInFileCount     = 0
)

type image struct {
	line int
	url  string
}

type downloadChannels struct {
	ok      chan string
	failure chan string
}

// DownloadFromTxtFile entry point for downloading images from a text file.
func DownloadFromTxtFile(cmd *cobra.Command, args []string) {
	startTime := time.Now()
	defer handlePanic()

	flags, err := flags.GetFlagValues(cmd)

	if err != nil {
		panic(err)
	}

	createOutDir(flags.OutputDir)
	images := extractImageURLsFromFile(flags.FilePath)
	defer displaySummary(startTime)

	okChan := make(chan string)
	failureChan := make(chan string, imagesInFileCount)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, img := range images {
		wg.Add(1)

		go func(img image) {
			defer wg.Done()
			downloadImage(downloadImageParams{
				img:       img,
				outputDir: flags.OutputDir,
				ch: downloadChannels{
					ok:      okChan,
					failure: failureChan,
				},
				mu:      &mu,
				timeout: flags.Timeout,
			})
		}(img)
	}

	go func() {
		for res := range okChan {
			fmt.Printf("[MESSAGE]: %s\n", res)
		}
	}()

	go func() {
		for error := range failureChan {
			fmt.Printf("[ERROR]: %s\n", error)
		}
	}()

	wg.Wait()
	close(okChan)
	close(failureChan)
}

func displaySummary(startTime time.Time) {
	fmt.Print("\n\n========== DOWNLOAD SUMMARY ==========\n\n\n")
	fmt.Printf("Images found in files: %d\n", fileLinesCount)
	fmt.Printf("Valid image URLs count: %d\n", imagesInFileCount)
	fmt.Printf("Images downloaded: %d of %d\n", downloadedImagesCount, imagesInFileCount)
	fmt.Printf("Images downloading took %s\n", time.Since(startTime))
}

type downloadImageParams struct {
	img       image
	outputDir string
	ch        downloadChannels
	mu        *sync.Mutex
	timeout   uint
}

func downloadImage(params downloadImageParams) {
	startTime := time.Now()
	params.ch.ok <- fmt.Sprintf("Downloading image in line %d (%s)", params.img.line, params.img.url)

	res, cancel := fetchImage(fetchImageParams{
		img:     params.img,
		failCh:  params.ch.failure,
		timeout: params.timeout,
	})

	if res == nil {
		return
	}

	defer res.Body.Close()
	defer cancel()

	if err := validateDownloadedImage(res, params.img); err != nil {
		params.ch.failure <- err.Error()
		return
	}

	filename := path.Base(params.img.url)
	destination := filepath.Join(params.outputDir, filename)
	file, err := os.Create(destination)

	if err != nil {
		params.ch.failure <- fmt.Sprintf("Error creating file for %s | %v", params.img.url, err)
		return
	}

	defer file.Close()

	if _, err = io.Copy(file, res.Body); err != nil {
		params.ch.failure <- fmt.Sprintf("Error saving image in line %d (%s) | %v", params.img.line, params.img.url, err)
		return
	}

	params.ch.ok <- fmt.Sprintf("Image in line %d (%s) downloaded successfully in %s!", params.img.line, params.img.url, time.Since(startTime))

	defer params.mu.Unlock()
	params.mu.Lock()
	downloadedImagesCount++
}

func validateDownloadedImage(res *http.Response, img image) error {
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download image in line %d (%s) | Status Code: %d", img.line, img.url, res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return fmt.Errorf("The resource downloaded in line %d (%s) is not an image | Content Type: %s", img.line, img.url, contentType)
	}

	return nil
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Printf("An error occurred: %v\n", r)
	}
}
