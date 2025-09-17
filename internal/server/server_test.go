package server

import (
	"booru-server/internal/config"
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockClient is a mock implementation of the booru.BooruClient interface.
type MockClient struct {
	NameFunc   func() string
	SearchFunc func(ctx context.Context, params booru.SearchParams) ([]booru.Image, error)
}

func (m *MockClient) Name() string {
	return m.NameFunc()
}

func (m *MockClient) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	return m.SearchFunc(ctx, params)
}

func TestServer_HandleSearch_Success(t *testing.T) {
	client1 := &MockClient{
		NameFunc: func() string { return "mock1" },
		SearchFunc: func(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
			return []booru.Image{{ID: "1", Provider: "mock1"}}, nil
		},
	}
	client2 := &MockClient{
		NameFunc: func() string { return "mock2" },
		SearchFunc: func(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
			return []booru.Image{{ID: "2", Provider: "mock2"}}, nil
		},
	}

	authCfg := config.AuthConfig{Enabled: false}
	rateLimitCfg := config.RateLimitConfig{Enabled: false}
	server := New([]booru.BooruClient{client1, client2}, authCfg, rateLimitCfg)
	ts := httptest.NewServer(server.router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/search")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", res.Status)
	}

	var images []booru.Image
	if err := json.NewDecoder(res.Body).Decode(&images); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(images) != 2 {
		t.Fatalf("Expected 2 images, got %d", len(images))
	}
}

func TestServer_HandleSearch_Auth(t *testing.T) {
	authCfg := config.AuthConfig{Enabled: true, Tokens: []string{"valid-token"}}
	rateLimitCfg := config.RateLimitConfig{Enabled: false}
	server := New([]booru.BooruClient{}, authCfg, rateLimitCfg)
	ts := httptest.NewServer(server.router)
	defer ts.Close()

	// No token
	res, _ := http.Get(ts.URL + "/api/search")
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", res.StatusCode)
	}

	// Invalid token
	req, _ := http.NewRequest("GET", ts.URL+"/api/search", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	res, _ = http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", res.StatusCode)
	}

	// Valid token
	req, _ = http.NewRequest("GET", ts.URL+"/api/search", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	res, _ = http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
}

func TestServer_HandleSearch_RateLimit(t *testing.T) {
	authCfg := config.AuthConfig{Enabled: false}
	rateLimitCfg := config.RateLimitConfig{Enabled: true, RequestsPerSecond: 1, Burst: 1}
	server := New([]booru.BooruClient{}, authCfg, rateLimitCfg)
	ts := httptest.NewServer(server.router)
	defer ts.Close()

	// First request should succeed
	res, _ := http.Get(ts.URL + "/api/search")
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	// Second request should be rate limited
	res, _ = http.Get(ts.URL + "/api/search")
	if res.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got %d", res.StatusCode)
	}
}
