package purrbotsite

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const clientName = "purrbot.site"

// Client is a client for the purrbot.site API.
type Client struct {
	BaseURL    string
	httpClient *http.Client
}

// New creates a new purrbot.site client.
func New() *Client {
	return &Client{
		BaseURL:    "https://purrbot.site/api/v2",
		httpClient: &http.Client{},
	}
}

// Name returns the name of the client.
func (c *Client) Name() string {
	return clientName
}

// Search queries the purrbot.site API.
func (c *Client) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	if len(params.Tags) == 0 {
		return nil, fmt.Errorf("at least one tag is required for purrbot.site search")
	}
	category := params.Tags[0]

	endpointType := "sfw"
	if params.NSFW != nil && *params.NSFW {
		endpointType = "nsfw"
	}

	// For now, we'll just get images, not gifs.
	imageType := "img"

	reqURL := fmt.Sprintf("%s/img/%s/%s/%s", c.BaseURL, endpointType, category, imageType)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	var apiResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResponse.Error {
		return nil, fmt.Errorf("API returned an error")
	}

	// The API returns a single image, so we wrap it in a slice
	return toBooruImages([]Response{apiResponse}, category, endpointType == "nsfw"), nil
}

func toBooruImages(responses []Response, category string, isNSFW bool) []booru.Image {
	booruImages := make([]booru.Image, len(responses))
	for i, res := range responses {
		booruImages[i] = booru.Image{
			ID:       res.Link,
			URL:      res.Link,
			Tags:     []string{category},
			NSFW:     isNSFW,
			Provider: clientName,
		}
	}
	return booruImages
}
