package booru

import "time"

// Image represents a single image from a booru source.
// It's a unified format for all API clients.
type Image struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Source    string    `json:"source,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Width     int       `json:"width,omitempty"`
	Height    int       `json:"height,omitempty"`
	Score     int       `json:"score,omitempty"`
	NSFW      bool      `json:"nsfw"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Provider  string    `json:"provider"` // e.g., "waifu.im", "waifu.pics"
}

// SearchParams holds all possible search parameters for querying booru APIs.
type SearchParams struct {
	Tags      []string
	NSFW      *bool // Use a pointer to distinguish between false and not set
	Width     *int
	Height    *int
	OrderBy   string
	Limit     int
}
