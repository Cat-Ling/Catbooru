package picre

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const clientName = "pic.re"

// Client is a client for the pic.re API.
type Client struct {
	BaseURL    string
	httpClient *http.Client
}

// New creates a new pic.re client.
func New() *Client {
	return &Client{
		BaseURL:    "https://pic.re",
		httpClient: &http.Client{},
	}
}

// Name returns the name of the client.
func (c *Client) Name() string {
	return clientName
}

// Search queries the pic.re API.
func (c *Client) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	reqURL, err := url.Parse(c.BaseURL + "/image.json")
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	q := reqURL.Query()
	if len(params.Tags) > 0 {
		q.Set("in", strings.Join(params.Tags, ","))
	}
	// pic.re does not have a direct NSFW filter, so we can't use params.NSFW
	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
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

	var apiResponse Image
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// The API returns a single image, so we wrap it in a slice
	return toBooruImages([]Image{apiResponse}), nil
}

func toBooruImages(images []Image) []booru.Image {
	booruImages := make([]booru.Image, len(images))
	for i, img := range images {
		booruImages[i] = booru.Image{
			ID:       strconv.Itoa(img.ID),
			URL:      img.FileURL,
			Source:   img.Source,
			Tags:     img.Tags,
			Width:    img.Width,
			Height:   img.Height,
			Provider: clientName,
		}
	}
	return booruImages
}
