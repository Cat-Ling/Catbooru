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

const moduleName = "pic.re"

var _ booru.BooruModule = (*Module)(nil)

// Module is a module for the pic.re API.
type Module struct {
	baseURL    string
	httpClient *http.Client
}

// Option is a functional option for configuring the pic.re module.
type Option func(*Module)

// WithBaseURL sets the base URL for the pic.re module.
func WithBaseURL(baseURL string) Option {
	return func(m *Module) {
		m.baseURL = baseURL
	}
}

// New creates a new pic.re module.
func New(opts ...Option) *Module {
	m := &Module{
		baseURL:    "https://pic.re",
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

// Search queries the pic.re API.
func (c *Module) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	reqURL, err := url.Parse(c.baseURL + "/image.json")
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
			Provider: moduleName,
		}
	}
	return booruImages
}
