package purrbotsite

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
		expectedPath := "/img/sfw/pat/img"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Link: "https://example.com/image.jpg",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := New(server.URL)
	isNsfw := false
	params := booru.SearchParams{Tags: []string{"pat"}, NSFW: &isNsfw}
	images, err := client.Search(context.Background(), params)

	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(images) != 1 {
		t.Fatalf("Expected 1 image, got %d", len(images))
	}

	if images[0].URL != "https://example.com/image.jpg" {
		t.Errorf("Expected image URL 'https://example.com/image.jpg', got '%s'", images[0].URL)
	}
}
