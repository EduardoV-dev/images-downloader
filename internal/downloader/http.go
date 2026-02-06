package downloader

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type fetchImageParams struct {
	img     image
	failCh  chan<- string
	timeout uint
}

func fetchImage(params fetchImageParams) (*http.Response, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(params.timeout)*time.Second)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, params.img.url, nil)
	if err != nil {
		params.failCh <- fmt.Sprintf("Image in line %d (%s) had an error | %v | timeout: %d seconds", params.img.line, params.img.url, err, params.timeout)
		return nil, cancel
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		params.failCh <- fmt.Sprintf("Could not download image in line %d (%s) | %v | timeout: %d seconds", params.img.line, params.img.url, err, params.timeout)
		return nil, cancel
	}

	return res, cancel
}
