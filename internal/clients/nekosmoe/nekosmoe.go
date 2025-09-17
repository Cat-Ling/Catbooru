package nekosmoe

import (
	"booru-server/pkg/booru"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	clientName = "nekos.moe"
	imageHost  = "https://nekos.moe/image"
)

// Client is a client for the nekos.moe API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// New creates a new nekos.moe client.
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

// Search queries the nekos.moe API.
func (c *Client) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	searchBody := map[string]interface{}{}
	if len(params.Tags) > 0 {
		searchBody["tags"] = params.Tags
	}
	if params.NSFW != nil {
		searchBody["nsfw"] = *params.NSFW
	}
	if params.Limit > 0 {
		searchBody["limit"] = params.Limit
	}

	body, err := json.Marshal(searchBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	reqURL := c.baseURL + "/images/search"
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

	var apiResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return toBooruImages(apiResponse.Images), nil
}

func toBooruImages(images []Image) []booru.Image {
	booruImages := make([]booru.Image, len(images))
	for i, img := range images {
		booruImages[i] = booru.Image{
			ID:        img.ID,
			URL:       fmt.Sprintf("%s/%s.jpg", imageHost, img.ID), // .jpg is a guess, might need adjustment
			Source:    fmt.Sprintf("https://nekos.moe/post/%s", img.ID),
			Tags:      img.Tags,
			Score:     img.Likes,
			NSFW:      img.Nsfw,
			CreatedAt: img.CreatedAt,
			Provider:  clientName,
		}
	}
	return booruImages
}
