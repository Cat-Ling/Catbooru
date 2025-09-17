package nekosapi

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const moduleName = "nekosapi.com"

var _ booru.BooruModule = (*Module)(nil)

// Module is a module for the nekosapi.com API.
type Module struct {
	baseURL    string
	httpClient *http.Client
}

// Option is a functional option for configuring the nekosapi.com module.
type Option func(*Module)

// WithBaseURL sets the base URL for the nekosapi.com module.
func WithBaseURL(baseURL string) Option {
	return func(m *Module) {
		m.baseURL = baseURL
	}
}

// New creates a new nekosapi.com module.
func New(opts ...Option) *Module {
	m := &Module{
		baseURL:    "https://api.nekosapi.com/v4",
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

// Search queries the nekosapi.com API.
func (c *Module) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	reqURL, err := url.Parse(c.baseURL + "/images/random")
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	q := reqURL.Query()
	if len(params.Tags) > 0 {
		q.Set("tags", strings.Join(params.Tags, ","))
	}
	if params.NSFW != nil {
		if *params.NSFW {
			q.Set("rating", "explicit,borderline")
		} else {
			q.Set("rating", "safe,suggestive")
		}
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

	var apiResponse []Image
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return toBooruImages(apiResponse), nil
}

func toBooruImages(images []Image) []booru.Image {
	booruImages := make([]booru.Image, len(images))
	for i, img := range images {
		tags := make([]string, len(img.Tags))
		for j, tag := range img.Tags {
			tags[j] = tag.Name
		}

		booruImages[i] = booru.Image{
			ID:        img.ID,
			URL:       img.FileURL,
			Source:    img.Source.URL,
			Tags:      tags,
			NSFW:      img.Rating == "explicit" || img.Rating == "borderline",
			CreatedAt: img.CreatedAt,
			Provider:  moduleName,
		}
	}
	return booruImages
}
