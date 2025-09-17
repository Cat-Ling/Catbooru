package waifupics

import (
	"booru-server/pkg/booru"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const clientName = "waifu.pics"

var _ booru.BooruClient = (*Client)(nil)

// Client is a client for the waifu.pics API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// New creates a new waifu.pics client.
func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// Name returns the name of the client.
func (c *Client) Name() string {
	return clientName
}

// Search queries the waifu.pics API.
func (c *Client) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	if len(params.Tags) == 0 {
		return nil, fmt.Errorf("at least one tag is required for waifu.pics search")
	}
	category := params.Tags[0]

	endpointType := "sfw"
	if params.NSFW != nil && *params.NSFW {
		endpointType = "nsfw"
	}

	reqURL := fmt.Sprintf("%s/many/%s/%s", c.baseURL, endpointType, category)

	// The API expects an empty JSON body for the "many" endpoint
	body, err := json.Marshal(request{Exclude: []string{}})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	var apiResponse response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return toBooruImages(apiResponse.Files, category, endpointType == "nsfw"), nil
}

func toBooruImages(urls []string, category string, isNSFW bool) []booru.Image {
	booruImages := make([]booru.Image, len(urls))
	for i, u := range urls {
		booruImages[i] = booru.Image{
			ID:       u, // waifu.pics doesn't have a stable ID, so we use the URL
			URL:      u,
			Tags:     []string{category},
			NSFW:     isNSFW,
			Provider: clientName,
		}
	}
	return booruImages
}
