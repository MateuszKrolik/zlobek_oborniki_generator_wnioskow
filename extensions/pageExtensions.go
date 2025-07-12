package extensions

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type Pages map[int][]byte

func (pages Pages) ToPdf(gotenbergBaseUrl string) ([]byte, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	var mergedHTML bytes.Buffer
	for i := 1; i <= len(pages); i++ {
		mergedHTML.Write(pages[i])
	}

	// gotenberg requires at least 1 file named "index.html"
	part, err := writer.CreateFormFile("files", "index.html")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := part.Write(mergedHTML.Bytes()); err != nil {
		return nil, fmt.Errorf("failed to write HTML content: %w", err)
	}

	writer.Close()

	gontenbergUrl := fmt.Sprintf("%v/forms/chromium/convert/html", gotenbergBaseUrl)

	resp, err := http.Post(
		gontenbergUrl,
		writer.FormDataContentType(),
		&body,
	)
	if err != nil {
		return nil, fmt.Errorf("Gotenberg request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Gotenberg error (%d): %s", resp.StatusCode, string(errorBody))
	}

	return io.ReadAll(resp.Body)
}
