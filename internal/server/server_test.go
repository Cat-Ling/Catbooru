package server

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockModule is a mock implementation of the booru.BooruModule interface.
type MockModule struct {
	NameFunc   func() string
	SearchFunc func(ctx context.Context, params booru.SearchParams) ([]booru.Image, error)
}

func (m *MockModule) Name() string {
	return m.NameFunc()
}

func (m *MockModule) Search(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
	return m.SearchFunc(ctx, params)
}

func TestServer_HandleSearch(t *testing.T) {
	module1 := &MockModule{
		NameFunc: func() string { return "mock1" },
		SearchFunc: func(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
			return []booru.Image{{ID: "1", Provider: "mock1"}}, nil
		},
	}
	module2 := &MockModule{
		NameFunc: func() string { return "mock2" },
		SearchFunc: func(ctx context.Context, params booru.SearchParams) ([]booru.Image, error) {
			return []booru.Image{{ID: "2", Provider: "mock2"}}, nil
		},
	}

	server := New([]booru.BooruModule{module1, module2}, "test-version")
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
