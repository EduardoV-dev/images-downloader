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

func extractImageURLsFromFile(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	imageUrls := []string{}

	for scanner.Scan() {
		url := scanner.Text()

		if ok := isValidUrl(url); ok {
			imageUrls = append(imageUrls, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return imageUrls
}

func isValidUrl(url string) bool {
	if err := isASingleImage(url); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func isASingleImage(url string) error {
	httpCount := strings.Count(url, "http")

	if httpCount > 1 {
		return fmt.Errorf("ERROR: %s contains multiple images on the same line, place each image on their own line\n", url)
	}

	return nil
}
