package purrbotsite

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const moduleName = "purrbot.site"

var _ booru.BooruModule = (*Module)(nil)

// Module is a module for the purrbot.site API.
type Module struct {
	baseURL    string
	httpClient *http.Client
}

// Option is a functional option for configuring the purrbot.site module.
type Option func(*Module)

// WithBaseURL sets the base URL for the purrbot.site module.
func WithBaseURL(baseURL string) Option {
	return func(m *Module) {
		m.baseURL = baseURL
	}
}

// New creates a new purrbot.site module.
func New(opts ...Option) *Module {
	m := &Module{
		baseURL:    "https://api.purrbot.site/v2",
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

// Search queries the purrbot.site API.
func (c *Module) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
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

	reqURL := fmt.Sprintf("%s/img/%s/%s/%s", c.baseURL, endpointType, category, imageType)
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

	return toBooruImage(apiResponse, category, endpointType == "nsfw"), nil
}

func toBooruImage(res Response, category string, isNSFW bool) []booru.Image {
	return []booru.Image{
		{
			ID:       res.Link,
			URL:      res.Link,
			Tags:     []string{category},
			NSFW:     isNSFW,
			Provider: moduleName,
		},
	}
}
