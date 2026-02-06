package downloader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createOutDir(outdirPath string) {
	if err := os.MkdirAll(outdirPath, 0755); err != nil {
		panic(err)
	}
}

func extractImageURLsFromFile(filePath string) []image {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	images := []image{}

	for scanner.Scan() {
		fileLinesCount++
		url := scanner.Text()

		if err := isValidUrl(url); err != nil {
			fmt.Printf("[ERROR]: %s\n", err.Error())
			continue
		}

		images = append(images, image{url: url, line: fileLinesCount})
		imagesInFileCount++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return images
}

func isValidUrl(url string) error {
	if err := isASingleImage(url); err != nil {
		return err
	}

	return nil
}

func isASingleImage(url string) error {
	httpCount := strings.Count(url, "http")

	if httpCount == 0 {
		return fmt.Errorf("%s is not a valid image url", url)
	}

	if httpCount > 1 {
		return fmt.Errorf("%s contains multiple images on the same line, place each image on their own line", url)
	}

	return nil
}
