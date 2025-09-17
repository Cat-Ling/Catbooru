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

const moduleName = "waifu.im"

var _ booru.BooruModule = (*Module)(nil)

// Module is a module for the waifu.im API.
type Module struct {
	baseURL    string
	httpClient *http.Client
}

// Option is a functional option for configuring the waifu.im module.
type Option func(*Module)

// WithBaseURL sets the base URL for the waifu.im module.
func WithBaseURL(baseURL string) Option {
	return func(m *Module) {
		m.baseURL = baseURL
	}
}

// New creates a new waifu.im module.
func New(opts ...Option) *Module {
	m := &Module{
		baseURL:    "https://api.waifu.im",
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// Name returns the name of the module.
func (c *Module) Name() string {
	return moduleName
}

// Search queries the waifu.im API.
func (c *Module) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
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
			Provider:  moduleName,
		}
	}
	return booruImages
}
