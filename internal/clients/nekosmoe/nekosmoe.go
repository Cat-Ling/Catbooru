package nekosmoe

import (
	"booru-server/pkg/booru"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const clientName = "nekos.moe"

var _ booru.BooruClient = (*Client)(nil)

// Client is a client for the nekos.moe API.
type Client struct {
	baseURL    string
	imageURL   string
	httpClient *http.Client
}

// New creates a new nekos.moe client.
func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		imageURL:   strings.Replace(baseURL, "/api/v1", "/image", 1),
		httpClient: &http.Client{},
	}
}

// Name returns the name of the client.
func (c *Client) Name() string {
	return clientName
}

// Search queries the nekos.moe API.
func (c *Client) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	searchReq := SearchRequest{
		Tags:  params.Tags,
		NSFW:  params.NSFW,
		Limit: params.Limit,
	}
	if params.OrderBy == "likes" {
		searchReq.Sort = "likes"
	}

	body, err := json.Marshal(searchReq)
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

	return c.toBooruImages(apiResponse.Images), nil
}

func (c *Client) toBooruImages(images []Image) []booru.Image {
	booruImages := make([]booru.Image, len(images))
	for i, img := range images {
		booruImages[i] = booru.Image{
			ID:        img.ID,
			URL:       fmt.Sprintf("%s/%s.png", c.imageURL, img.ID),
			Source:    fmt.Sprintf("%s/post/%s", strings.Replace(c.baseURL, "/api/v1", "", 1), img.ID),
			Tags:      img.Tags,
			Score:     img.Likes,
			NSFW:      img.Nsfw,
			CreatedAt: img.CreatedAt,
			Provider:  clientName,
		}
	}
	return booruImages
}
