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
