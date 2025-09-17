package waifupics

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Search(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/many/sfw/waifu" {
			t.Errorf("Expected to request '/many/sfw/waifu', got: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			Files []string `json:"files"`
		}{
			Files: []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := New(server.URL)
	isNsfw := false
	params := booru.SearchParams{Tags: []string{"waifu"}, NSFW: &isNsfw}
	images, err := client.Search(context.Background(), params)

	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(images) != 2 {
		t.Fatalf("Expected 2 images, got %d", len(images))
	}

	if images[0].URL != "https://example.com/image1.jpg" {
		t.Errorf("Expected image URL 'https://example.com/image1.jpg', got '%s'", images[0].URL)
	}
}
