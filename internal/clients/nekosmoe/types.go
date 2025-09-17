package nekosmoe

import "time"

// Response is the top-level structure of the nekos.moe API response.
type Response struct {
	Images []Image `json:"images"`
}

// Image is the structure of an image object from the nekos.moe API.
type Image struct {
	ID        string    `json:"id"`
	Tags      []string  `json:"tags"`
	Artist    string    `json:"artist"`
	Likes     int       `json:"likes"`
	Nsfw      bool      `json:"nsfw"`
	CreatedAt time.Time `json:"createdAt"`
}

// SearchRequest is the request body for the nekos.moe image search API.
type SearchRequest struct {
	Tags  []string `json:"tags,omitempty"`
	NSFW  *bool    `json:"nsfw,omitempty"`
	Limit int      `json:"limit,omitempty"`
	Sort  string   `json:"sort,omitempty"`
}
