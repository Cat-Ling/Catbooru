package waifuim

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

const clientName = "waifu.im"

// Client is a client for the waifu.im API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// New creates a new waifu.im client.
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

// Search queries the waifu.im API.
func (c *Client) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	reqURL, err := url.Parse(c.baseURL + "/search")
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	q := reqURL.Query()
	if len(params.Tags) > 0 {
		q.Set("included_tags", strings.Join(params.Tags, ","))
	}
	if params.NSFW != nil {
		q.Set("is_nsfw", strconv.FormatBool(*params.NSFW))
	}
	if params.Limit > 0 {
		q.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.Width != nil {
		q.Set("width", fmt.Sprintf(">=%d", *params.Width))
	}
	if params.Height != nil {
		q.Set("height", fmt.Sprintf(">=%d", *params.Height))
	}
	if params.OrderBy != "" {
		q.Set("order_by", params.OrderBy)
	}
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

	var apiResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return toBooruImages(apiResponse.Images), nil
}

func toBooruImages(images []Image) []booru.Image {
	booruImages := make([]booru.Image, len(images))
	for i, img := range images {
		tags := make([]string, len(img.Tags))
		for j, tag := range img.Tags {
			tags[j] = tag.Name
		}

		booruImages[i] = booru.Image{
			ID:        strconv.Itoa(img.ImageID),
			URL:       img.URL,
			Source:    img.Source,
			Tags:      tags,
			Width:     img.Width,
			Height:    img.Height,
			Score:     img.Favorites,
			NSFW:      img.IsNSFW,
			CreatedAt: img.UploadedAt,
			Provider:  clientName,
		}
	}
	return booruImages
}
